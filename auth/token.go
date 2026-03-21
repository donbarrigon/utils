package auth

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"github.com/donbarrigon/utils/herror"
	"github.com/donbarrigon/utils/logs"
)

// genera un token exadecimal de 32 bytes
func GenerateHexToken() (string, herror.Error) {
	bytes := make([]byte, 32)
	if _, e := rand.Read(bytes); e != nil {
		logs.Error("Failed to generate secure token")
		return "", herror.InternalServerErrorMsg(e, "Failed to generate secure token")
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateVerificationCode() (string, herror.Error) {
	code := ""
	for range 6 {
		num, e := rand.Int(rand.Reader, big.NewInt(10))
		if e != nil {
			return "", herror.InternalServerErrorMsg(e, "Failed to generate secure token")
		}
		code += num.String()
	}
	return code, nil
}
