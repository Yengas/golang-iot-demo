package tokenizer

import (
	"github.com/dgrijalva/jwt-go"
	"iot-demo/pkg/auth"
	"strconv"
)

type DeviceJWT AuthJWT

func (dj DeviceJWT) Create(cred *auth.DeviceCredential) (auth.Token, error) {
	claims := jwt.MapClaims{
		"id": strconv.Itoa(cred.DeviceID),
	}

	token, err := AuthJWT(dj).Sign(claims)
	if err != nil {
		return "", nil
	}

	return auth.Token(token), nil
}
