package main

import (
	"backend/lib/config"
	"backend/lib/helper"
	"backend/lib/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type LoginPayload struct {
	Address       string `json:"address"`
	Signature     string `json:"signature"`
	OriginMessage string `json:"origin_message"`
}

type LoginResponseBody struct {
	Jwt string `json:"jwt"`
}

func login(c echo.Context) error {
	defer c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	lp := new(LoginPayload)
	if err := c.Bind(lp); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "request.format.error",
		})
	}

	body, err1 := process(lp)
	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err1.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": body,
	})
}

func process(loginPayload *LoginPayload) (*LoginResponseBody, error) {
	verifed := helper.VerifyEip191Sig(loginPayload.Address, loginPayload.OriginMessage, loginPayload.Signature)
	if !verifed {
		return nil, errors.New("request.verify.error")
	}
	address := strings.ToLower(loginPayload.Address)

	_, err1 := models.FindOrCreateUser(env_config.FaunadbClient, address)
	if err1 != nil {
		return nil, errors.New("request.db.error")
	}

	claims := jwt.MapClaims{"address": address}
	jwtStr, err2 := helper.EncodeJwt(claims, env_config.Conf.JwtRsaPrivateKeyPem, int64(86400))
	if err2 != nil {
		return nil, errors.New("request.jwt.error")
	}
	return &LoginResponseBody{Jwt: jwtStr}, nil
}
