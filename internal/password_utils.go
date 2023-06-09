package internal

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"

	"project-ricotta/bechamel-api/config"
)

func HashPassword(password string) string {
	salt := config.RuntimeConfig.PasswordSalt()
	outKey, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(outKey)
}
