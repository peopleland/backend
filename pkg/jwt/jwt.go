package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type MapClaims struct {
	jwt.MapClaims
}

func NewMapClaims(userid, address string) *MapClaims {
	return &MapClaims{map[string]interface{}{"userid": userid, "address": address}}
}

func EncodeJwt(claims *MapClaims, privateKeyPemStr string, exp int64) (string, error) {
	now := time.Now().Unix()
	claims.MapClaims["iat"] = now
	claims.MapClaims["nbf"] = now
	claims.MapClaims["exp"] = now + exp

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPemStr))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims.MapClaims)

	tokenString, err2 := token.SignedString(privateKey)
	if err2 != nil {
		return "", err2
	}
	return tokenString, nil
}

func DecodeJwt(jwtStr string, publicKeyStr string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(jwtStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
		if err != nil {
			return "", err
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	} else {
		return nil, err
	}
}
