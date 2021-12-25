package main

import (
	env_config "backend/lib/config"
	"backend/lib/helper"
	"backend/lib/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func changeProfile(c echo.Context) error {
	address := helper.GetCurrentUserAddress(c)
	if address == "" {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "request.user.not_found",
		})
	}

	updateData := map[string]string{}
	err := json.NewDecoder(c.Request().Body).Decode(&updateData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "request.body.format.error",
		})
	}

	filterUpdateData := map[string]interface{}{}
	for key, value := range updateData {
		if key == "name" || key == "twitter" {
			filterUpdateData[key] = value
		}
	}

	user, err1 := models.UpdateUserByAddress(env_config.FaunadbClient, address, filterUpdateData)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "request.db.error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": user.Data,
	})
}
