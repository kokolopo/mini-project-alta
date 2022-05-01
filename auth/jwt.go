package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte("fahmi_A12019")

type Service interface {
	GenerateTokenJWT(id int, name string) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokenJWT(userId int, name string) (string, error) {
	claim := jwt.MapClaims{}
	claim["id"] = userId
	claim["name"] = name
	claim["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
