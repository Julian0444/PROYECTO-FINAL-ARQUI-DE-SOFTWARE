package utils

import "golang.org/x/crypto/bcrypt"

type Crypto struct{}

type IHash interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}

func NewCrypto() *Crypto {
	return &Crypto{}
}

func (c *Crypto) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (c *Crypto) ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil

}
