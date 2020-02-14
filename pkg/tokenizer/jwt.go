package tokenizer

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Config struct {
	Secret []byte
}

type AuthJWT struct {
	config Config
}

func (aj AuthJWT) Sign(claim jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	str, err := token.SignedString(aj.config.Secret)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (aj AuthJWT) Parse(str string) (*jwt.Token, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return aj.config.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func NewJWT(config Config) AuthJWT {
	return AuthJWT{config: config}
}
