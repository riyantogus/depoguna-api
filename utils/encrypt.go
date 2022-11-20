package utils

import "golang.org/x/crypto/bcrypt"

type EncryptUtil interface {
	HashAndSalt(pwd []byte) string
	VerifyPassword(hashedPwd string, plainPwd []byte) bool
}

type encryptUtil struct{}

func NewEncryptUtil() EncryptUtil {
	return &encryptUtil{}
}

func (e *encryptUtil) HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash a password.")
	}
	return string(hash)
}

func (e *encryptUtil) VerifyPassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	if err := bcrypt.CompareHashAndPassword(byteHash, plainPwd); err != nil {
		return false
	}
	return true
}
