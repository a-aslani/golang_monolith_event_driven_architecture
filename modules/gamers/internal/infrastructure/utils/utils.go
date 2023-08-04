package utils

import (
	"github.com/a-aslani/golang_monolith_event_driven_architecture/modules/gamers/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Utils struct{}

var _ domain.Utils = (*Utils)(nil)

func NewUtils() Utils {
	return Utils{}
}

func (u Utils) Hash(str string) (string, error) {
	strByte := []byte(str)
	hashed, err := bcrypt.GenerateFromPassword(strByte, bcrypt.DefaultCost)
	return string(hashed), err
}

func (u Utils) CompareHash(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
