package http

import (
	"backend/pkg/http/binding"
	"context"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Ctx context.Context
}

func (c *Context) BindQuery(v interface{}) error {
	return binding.BindQuery(c.QueryParams(), v)
}
