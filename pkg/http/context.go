package http

import (
	"backend/pkg/http/binding"
	"context"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	Ctx context.Context
}

func (c *Context) Vars() url.Values {
	raws := mux.Vars(c.Context.Request())
	vars := make(url.Values, len(raws))
	for k, v := range raws {
		vars[k] = []string{v}
	}
	return vars
}

func (c *Context) BindQuery(v interface{}) error {
	return binding.BindQuery(c.Vars(), v)
}
