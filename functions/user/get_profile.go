package main

import (
	env_config "backend/lib/config"
	"backend/lib/helper"
	"backend/lib/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getProfile(c echo.Context) error {
	jwtStr := c.Request().Header.Get("authorization")
	jwtMap, err := helper.DecodeJwt(jwtStr, env_config.Conf.JwtRsaPublicKeyPem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "request.jwt.format.error",
		})
	}
	address := (*jwtMap)["address"].(string)
	user, err1 := models.GetOneUserByAddress(env_config.FaunadbClient, address)

	if err1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "request.user.not_found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user.Data,
	})
}
