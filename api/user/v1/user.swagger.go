package v1

import (
	"backend/pkg/http"
	"context"
)

type UserLambdaServer interface {
	Login(context.Context, *LoginPayLoad) (*LoginResponse, error)
	GetProfile(context.Context) (*UserProfile, error)
	PutProfile(context.Context, *PutProfilePayLoad) (*UserProfile, error)
	ConnectTwitter(context.Context, *ConnectTwitterPayLoad) (*UserProfile, error)
}

func userLambdaLoginHandler(serv UserLambdaServer) func(http.Context) error {
	return func(ctx http.Context) error {
		var in LoginPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		out, err := serv.Login(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		reply := out

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": reply,
		})
	}
}

func userLambdaGetProfileHandler(serv UserLambdaServer) func(http.Context) error {
	return func(ctx http.Context) error {
		out, err := serv.GetProfile(ctx.Ctx)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		reply := out

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": reply,
		})
	}
}

func userLambdaPutProfileHandler(serv UserLambdaServer) func(http.Context) error {
	return func(ctx http.Context) error {
		var in PutProfilePayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		out, err := serv.PutProfile(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		reply := out

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": reply,
		})
	}
}

func userLambdaConnectTwitterHandler(serv UserLambdaServer) func(http.Context) error {
	return func(ctx http.Context) error {
		var in ConnectTwitterPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		out, err := serv.ConnectTwitter(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}

		reply := out

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": reply,
		})
	}
}

func RegisterUserLambdaServer(s *http.Server, serv UserLambdaServer) {
	g := s.GroupX("/.netlify/functions/user")

	g.POSTX("/login", userLambdaLoginHandler(serv))
	g.GETX("/profile", userLambdaGetProfileHandler(serv))
	g.PUTX("/profile", userLambdaPutProfileHandler(serv))
	g.PUTX("/connect_twitter", userLambdaConnectTwitterHandler(serv))
}
