package auth

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/donbarrigon/utils/config"
	"github.com/donbarrigon/utils/herror"
	"github.com/vmihailenco/msgpack/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session struct {
	//ID          bson.ObjectID
	Token       string
	User        SessionUser
	IP          string
	Agent       string
	Lang        string
	Fingerprint string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ExpiresAt   time.Time
	writer      http.ResponseWriter
	request     *http.Request
}

var muSession = sync.Map{}

// ================================================================
//            FUNCIONES PARA LA INTERFAZ handler.Auth
// ================================================================

func (s *Session) Can(permission string) herror.Error {
	if s.User.Can(permission) {
		return nil
	}
	return herror.Forbidden(nil)
}

func (s *Session) HasRole(role string) herror.Error {
	if s.User.HasRole(role) {
		return nil
	}
	return herror.Forbidden(nil)
}

func (s *Session) UserID() bson.ObjectID {
	return s.User.GetID()
}

// ================================================================
//                  FUNCIONES AUXILIARES
// ================================================================

func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

func (s *Session) Save() herror.Error {
	echan := make(chan error, 2)
	go func() {
		muToken := s.muToken()
		muToken.Lock()
		defer muToken.Unlock()
		echan <- s.saveFileSession()
	}()
	go func() {
		muUser := s.muUser()
		muUser.Lock()
		defer muUser.Unlock()
		echan <- s.addFileUserIndex()
	}()

	if e := <-echan; e != nil {
		return herror.InternalServerError(e)
	}
	if e := <-echan; e != nil {
		return herror.InternalServerError(e)
	}
	return nil
}

func (s *Session) Refresh() herror.Error {
	muToken := s.muToken()
	muToken.Lock()
	defer muToken.Unlock()
	s.UpdatedAt = time.Now()
	s.ExpiresAt = expiresAt()
	return s.saveFileSession()
}

func (s *Session) Destroy() herror.Error {
	muToken := s.muToken()
	muToken.Lock()
	defer muToken.Unlock()
	if e := s.deleteFileSession(); e != nil {
		return e
	}
	s.ClearCookie()

	muUser := s.muUser()
	muUser.Lock()
	defer muUser.Unlock()
	return s.removeFileUserIndex()
}

func (s *Session) SetCookie() {
	http.SetCookie(s.writer, &http.Cookie{
		Name:     "session",
		Value:    s.Token,
		Path:     "/",
		HttpOnly: true,
		Secure:   config.ServerHttpsEnabled,
		SameSite: http.SameSiteLaxMode,
		Expires:  s.ExpiresAt,
	})
}

func (s *Session) ClearCookie() {
	http.SetCookie(s.writer, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   config.ServerHttpsEnabled,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
	})
}

func (s *Session) saveFileSession() herror.Error {
	path, filename := fileSession(s.Token)
	encoded, e := msgpack.Marshal(s)
	if e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to encode session data")
	}
	if e := os.MkdirAll(path, 0755); e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to create session directory")
	}
	if e := os.WriteFile(path+filename, encoded, 0644); e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to write session file")
	}
	return nil
}

func (s *Session) saveFileUserIndex(data map[string]time.Time) herror.Error {
	path, filename := fileUserIndex(s.User.GetID().Hex())
	encoded, e := msgpack.Marshal(data)
	if e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to encode session index")
	}
	if e := os.MkdirAll(path, 0755); e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to create session index directory")
	}
	if e := os.WriteFile(path+filename, encoded, 0644); e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to write session index file")
	}
	return nil
}

func (s *Session) addFileUserIndex() herror.Error {
	data, he := s.readFileUserIndex()
	if he != nil {
		return he
	}
	data[s.Token] = s.CreatedAt
	return s.saveFileUserIndex(data)
}

func (s *Session) deleteFileSession() herror.Error {
	path, fileName := fileSession(s.Token)
	if e := os.Remove(path + fileName); e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to delete session file")
	}
	return nil
}

func (s *Session) removeFileUserIndex() herror.Error {
	data, he := s.readFileUserIndex()
	if he != nil {
		return he
	}
	delete(data, s.Token)
	if len(data) == 0 {
		return s.deleteFileUserIndex()
	}
	return s.saveFileUserIndex(data)
}

func (s *Session) deleteFileUserIndex() herror.Error {
	path, fileName := fileUserIndex(s.User.GetID().Hex())
	if e := os.Remove(path + fileName); e != nil {
		return herror.InternalServerErrorMsg(e, "Failed to delete session index file")
	}
	return nil
}

func (s *Session) readFileUserIndex() (map[string]time.Time, herror.Error) {
	data := map[string]time.Time{}
	path, filename := fileUserIndex(s.User.GetID().Hex())
	info, e := os.Stat(path + filename)
	if e == nil && !info.IsDir() {
		encoded, e := os.ReadFile(path + filename)
		if e != nil {
			return nil, herror.InternalServerErrorMsg(e, "Failed to read session index file")
		}
		if e := msgpack.Unmarshal(encoded, &data); e != nil {
			return nil, herror.InternalServerErrorMsg(e, "Failed to decode session index data")
		}
	}
	return data, nil
}

func (s *Session) muUser() *sync.Mutex {
	mu, _ := muSession.LoadOrStore(s.User.GetID().Hex(), &sync.Mutex{})
	return mu.(*sync.Mutex)
}

func (s *Session) muToken() *sync.Mutex {
	mu, _ := muSession.LoadOrStore(s.Token, &sync.Mutex{})
	return mu.(*sync.Mutex)
}
