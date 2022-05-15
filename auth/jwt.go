package auth

import (
	"errors"
	"order_kafe/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateTokenJWT(id int, name string, role string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(config.InitConfiguration().JWT_KEY), nil

	})

	if err != nil {
		return token, err
	}

	return token, nil
}
