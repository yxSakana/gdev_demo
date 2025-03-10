package utility

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func MustGeneratePassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Panicf("error password hash: %v", err)
	}
	return base64.StdEncoding.EncodeToString(hash)
}

func CheckPassword(pwd, hash string) bool {
	b, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(b, []byte(pwd))
	return err == nil
}
