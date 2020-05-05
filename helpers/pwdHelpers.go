package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword from string
func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	return string(bytes), err
}

//CheckPasswordHash compares regular password to hashpassword. Returns true in success or false on failure
func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
