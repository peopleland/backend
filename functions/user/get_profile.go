package main

import (
	env_config "backend/lib/config"
	"backend/lib/helper"
	"backend/lib/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getProfile(c echo.Context) error {
	address := helper.GetCurrentUserAddress(c)
	if address == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "request.user.not_found",
		})
	}
	user, err1 := models.GetOneUserByAddress(env_config.FaunadbClient, address)

	if err1 != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "request.user.not_found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user.Data,
	})
}
