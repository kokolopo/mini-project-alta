package auth

import (
	"order_kafe/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateTokenJWT(id int, name string, role string) (string, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateTokenJWT(id int, name string, role string) (string, error) {

	claim := jwt.MapClaims{}
	claim["id"] = id
	claim["fullname"] = name
	claim["role"] = role
	claim["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(config.InitConfiguration().JWT_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
