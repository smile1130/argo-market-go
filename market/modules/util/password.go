package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func PasswordHashV1(username string, password string) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%s", username, password)))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
