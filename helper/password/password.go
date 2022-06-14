package password

import "golang.org/x/crypto/bcrypt"

func Bcrypt(password string) string {
	var bytes, _ = bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
