package auth

import (
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/donbarrigon/utils/config"
	"github.com/donbarrigon/utils/herror"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type SessionUser interface {
	GetID() bson.ObjectID
	Can(permission string) bool
	HasRole(role string) bool
}

func SessionStart(w http.ResponseWriter, r *http.Request, user SessionUser) (*Session, herror.Error) {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	tk, e := GenerateHexToken()
	if e != nil {
		return nil, e
	}
	s := &Session{
		//ID:          bson.NewObjectID(),
		Token:       tk,
		User:        user,
		IP:          host,
		Agent:       r.Header.Get("user-agent"),
		Lang:        r.Header.Get("accept-language"),
		Fingerprint: r.Header.Get("x-fingerprint"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ExpiresAt:   expiresAt(),
		writer:      w,
		request:     r,
	}
	s.SetCookie()
	if e := s.Save(); e != nil {
		return nil, e
	}
	return s, nil
}

func expiresAt() time.Time {
	return time.Now().Add(time.Duration(config.SessionLifetime) * time.Minute)
}

func GetSession(w http.ResponseWriter, r *http.Request) (*Session, herror.Error) {
	var token string
	cookie, err := r.Cookie("session")
	if err != nil {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			token = strings.TrimSpace(strings.TrimPrefix(authHeader, "bearer "))
		}

		if token == "" {
			return nil, herror.Unauthorized(err)
		}
	} else {
		token = cookie.Value
	}

	s, e := GetSessionByToken(token)
	if e != nil {
		return nil, e
	}
	s.writer = w
	s.request = r

	s.SetCookie()
	if e := s.Refresh(); e != nil {
		return nil, e
	}
	return s, nil
}

func GetSessionByToken(token string) (*Session, herror.Error) {
	s := &Session{}
	path, filename := fileSession(token)
	info, e := os.Stat(path + filename)
	if e == nil && !info.IsDir() {
		encoded, e := os.ReadFile(path + filename)
		if e != nil {
			return nil, herror.ForbiddenMsg(e, "Session not found or has expired")
		}
		if e := msgpack.Unmarshal(encoded, &s); e != nil {
			return nil, herror.ForbiddenMsg(e, "Session data is invalid or corrupted")
		}
	}
	return s, nil
}

func GetSessionsByUserID(hex string) ([]*Session, herror.Error) {
	sessions := []*Session{}
	tokens := map[string]time.Time{}
	path, filename := fileUserIndex(hex)
	info, e := os.Stat(path + filename)
	if e == nil && !info.IsDir() {
		encoded, e := os.ReadFile(path + filename)
		if e != nil {
			return nil, herror.ForbiddenMsg(e, "Failed to read user session index")
		}
		if e := msgpack.Unmarshal(encoded, &tokens); e != nil {
			return nil, herror.ForbiddenMsg(e, "User session index is invalid or corrupted")
		}
	}

	for token := range tokens {
		s, e := GetSessionByToken(token)
		if e != nil {
			return nil, e
		}
		sessions = append(sessions, s)
	}
	return sessions, nil
}

func fileSession(token string) (string, string) {
	return "tmp/sessions/" + token[:3] + "/" + token[3:6] + "/", token[6:]
}

func fileUserIndex(hex string) (string, string) {
	return "tmp/sessions/index/" + hex[:4] + "/" + hex[4:8] + "/", hex[8:]
}
