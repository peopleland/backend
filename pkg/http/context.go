package http

import (
	"context"

	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Ctx context.Context
}
