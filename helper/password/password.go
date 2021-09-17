package password

import (
	"github.com/devgek/webskeleton/helper/common"
	"golang.org/x/crypto/bcrypt"
)

//EncryptPassword create hashed password
func EncryptPassword(password string) string {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	common.PanicOnError(err)

	return string(encryptedPass)
}

//ComparePassword compare hashed password and possible plaintext equivalent
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
