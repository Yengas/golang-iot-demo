package jwt

import (
	"errors"
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

func (dj DeviceJWT) Parse(authToken auth.Token) (*auth.DeviceCredential, error) {
	jwtToken, err := AuthJWT(dj).Parse(string(authToken))
	if err != nil {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		deviceIDSTR := claims["id"].(string)
		deviceID, err := strconv.Atoi(deviceIDSTR)
		if err != nil {
			return nil, errors.New("could not parse the claim")
		}
		return &auth.DeviceCredential{DeviceID: deviceID}, nil
	}

	return nil, errors.New("no claims or token is not valid")
}
