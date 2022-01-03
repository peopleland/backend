// Code generated by protoc-gen-go-lambda. DO NOT EDIT.
// versions:
// protoc-gen-go-lambda 0.1

package v1

import (
	http "backend/pkg/http"
	binding "backend/pkg/http/binding"
	context "context"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.Version

type UserLambdaServer interface {
	ConnectDiscord(context.Context, *ConnectDiscordPayLoad) (*ConnectDiscordResponse, error)
	ConnectTelegram(context.Context, *ConnectTelegramPayLoad) (*ConnectTelegramResponse, error)
	ConnectTwitter(context.Context, *ConnectTwitterPayLoad) (*UserProfile, error)
	DisconnectSocial(context.Context, *DisconnectSocialPayLoad) (*DisconnectSocialResponse, error)
	GenVerifyCode(context.Context, *GenVerifyCodePayLoad) (*GenVerifyCodeResponse, error)
	GetProfile(context.Context, *GetProfilePayLoad) (*UserProfile, error)
	Login(context.Context, *LoginPayLoad) (*LoginResponse, error)
	PutProfile(context.Context, *PutProfilePayLoad) (*UserProfile, error)
	TelegramBotDMWebhooks(context.Context, *TelegramBotDMWebhooksPayLoad) (*TelegramBotDMWebhooksResponse, error)
}

func RegisterUserLambdaServer(s *http.Server, srv UserLambdaServer) {
	g := s.GroupX("/.netlify/functions")
	g.POSTX("/user/v1/login", _User_Login0_Lambda_Handler(srv))
	g.GETX("/user/v1/profile", _User_GetProfile0_Lambda_Handler(srv))
	g.PUTX("/user/v1/profile", _User_PutProfile0_Lambda_Handler(srv))
	g.PUTX("/user/v1/connect/twitter", _User_ConnectTwitter0_Lambda_Handler(srv))
	g.PUTX("/user/v1/connect/telegram", _User_ConnectTelegram0_Lambda_Handler(srv))
	g.PUTX("/user/v1/connect/discord", _User_ConnectDiscord0_Lambda_Handler(srv))
	g.POSTX("/user/v1/telegram/webhooks/dm", _User_TelegramBotDMWebhooks0_Lambda_Handler(srv))
	g.PUTX("/user/v1/gen_verify_code", _User_GenVerifyCode0_Lambda_Handler(srv))
	g.POSTX("/user/v1/social/disconnect", _User_DisconnectSocial0_Lambda_Handler(srv))
}

func _User_Login0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/Login")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.Login(ctx, req.(*LoginPayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.Login(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*LoginResponse)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_GetProfile0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfilePayLoad
		if err := ctx.BindQuery(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/GetProfile")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.GetProfile(ctx, req.(*GetProfilePayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.GetProfile(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*UserProfile)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_PutProfile0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PutProfilePayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/PutProfile")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.PutProfile(ctx, req.(*PutProfilePayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.PutProfile(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*UserProfile)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_ConnectTwitter0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConnectTwitterPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/ConnectTwitter")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.ConnectTwitter(ctx, req.(*ConnectTwitterPayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.ConnectTwitter(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*UserProfile)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_ConnectTelegram0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConnectTelegramPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/ConnectTelegram")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.ConnectTelegram(ctx, req.(*ConnectTelegramPayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.ConnectTelegram(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*ConnectTelegramResponse)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_ConnectDiscord0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConnectDiscordPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/ConnectDiscord")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.ConnectDiscord(ctx, req.(*ConnectDiscordPayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.ConnectDiscord(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*ConnectDiscordResponse)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_TelegramBotDMWebhooks0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TelegramBotDMWebhooksPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/TelegramBotDMWebhooks")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.TelegramBotDMWebhooks(ctx, req.(*TelegramBotDMWebhooksPayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.TelegramBotDMWebhooks(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*TelegramBotDMWebhooksResponse)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_GenVerifyCode0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenVerifyCodePayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/GenVerifyCode")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.GenVerifyCode(ctx, req.(*GenVerifyCodePayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.GenVerifyCode(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*GenVerifyCodeResponse)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

func _User_DisconnectSocial0_Lambda_Handler(srv UserLambdaServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DisconnectSocialPayLoad
		if err := ctx.Bind(&in); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//http.SetOperation(ctx,"/api.user.v1.User/DisconnectSocial")
		//h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		//	return srv.DisconnectSocial(ctx, req.(*DisconnectSocialPayLoad))
		//})
		//out, err := h(ctx, &in)
		out, err := srv.DisconnectSocial(ctx.Ctx, &in)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
		//reply := out.(*DisconnectSocialResponse)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"data": out,
		})
	}
}

//
//type UserLambdaClient interface {
//
//	ConnectDiscord(ctx context.Context, req *ConnectDiscordPayLoad, opts ...http.CallOption) (rsp *ConnectDiscordResponse, err error)
//
//	ConnectTelegram(ctx context.Context, req *ConnectTelegramPayLoad, opts ...http.CallOption) (rsp *ConnectTelegramResponse, err error)
//
//	ConnectTwitter(ctx context.Context, req *ConnectTwitterPayLoad, opts ...http.CallOption) (rsp *UserProfile, err error)
//
//	DisconnectSocial(ctx context.Context, req *DisconnectSocialPayLoad, opts ...http.CallOption) (rsp *DisconnectSocialResponse, err error)
//
//	GenVerifyCode(ctx context.Context, req *GenVerifyCodePayLoad, opts ...http.CallOption) (rsp *GenVerifyCodeResponse, err error)
//
//	GetProfile(ctx context.Context, req *GetProfilePayLoad, opts ...http.CallOption) (rsp *UserProfile, err error)
//
//	Login(ctx context.Context, req *LoginPayLoad, opts ...http.CallOption) (rsp *LoginResponse, err error)
//
//	PutProfile(ctx context.Context, req *PutProfilePayLoad, opts ...http.CallOption) (rsp *UserProfile, err error)
//
//	TelegramBotDMWebhooks(ctx context.Context, req *TelegramBotDMWebhooksPayLoad, opts ...http.CallOption) (rsp *TelegramBotDMWebhooksResponse, err error)
//
//}
//
//type UserLambdaClientImpl struct{
//	cc *http.Client
//}
//
//func NewUserLambdaClient (client *http.Client) UserLambdaClient {
//	return &UserLambdaClientImpl{client}
//}
//
//
//func (c *UserLambdaClientImpl) ConnectDiscord(ctx context.Context, in *ConnectDiscordPayLoad, opts ...http.CallOption) (*ConnectDiscordResponse, error) {
//	var out ConnectDiscordResponse
//	pattern := "/user/v1/connect/discord"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/ConnectDiscord"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) ConnectTelegram(ctx context.Context, in *ConnectTelegramPayLoad, opts ...http.CallOption) (*ConnectTelegramResponse, error) {
//	var out ConnectTelegramResponse
//	pattern := "/user/v1/connect/telegram"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/ConnectTelegram"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) ConnectTwitter(ctx context.Context, in *ConnectTwitterPayLoad, opts ...http.CallOption) (*UserProfile, error) {
//	var out UserProfile
//	pattern := "/user/v1/connect/twitter"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/ConnectTwitter"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) DisconnectSocial(ctx context.Context, in *DisconnectSocialPayLoad, opts ...http.CallOption) (*DisconnectSocialResponse, error) {
//	var out DisconnectSocialResponse
//	pattern := "/user/v1/social/disconnect"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/DisconnectSocial"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) GenVerifyCode(ctx context.Context, in *GenVerifyCodePayLoad, opts ...http.CallOption) (*GenVerifyCodeResponse, error) {
//	var out GenVerifyCodeResponse
//	pattern := "/user/v1/gen_verify_code"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/GenVerifyCode"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) GetProfile(ctx context.Context, in *GetProfilePayLoad, opts ...http.CallOption) (*UserProfile, error) {
//	var out UserProfile
//	pattern := "/user/v1/profile"
//	path := binding.EncodeURL(pattern, in, true)
//	opts = append(opts, http.Operation("/api.user.v1.User/GetProfile"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) Login(ctx context.Context, in *LoginPayLoad, opts ...http.CallOption) (*LoginResponse, error) {
//	var out LoginResponse
//	pattern := "/user/v1/login"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/Login"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) PutProfile(ctx context.Context, in *PutProfilePayLoad, opts ...http.CallOption) (*UserProfile, error) {
//	var out UserProfile
//	pattern := "/user/v1/profile"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/PutProfile"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
//func (c *UserLambdaClientImpl) TelegramBotDMWebhooks(ctx context.Context, in *TelegramBotDMWebhooksPayLoad, opts ...http.CallOption) (*TelegramBotDMWebhooksResponse, error) {
//	var out TelegramBotDMWebhooksResponse
//	pattern := "/user/v1/telegram/webhooks/dm"
//	path := binding.EncodeURL(pattern, in, false)
//	opts = append(opts, http.Operation("/api.user.v1.User/TelegramBotDMWebhooks"))
//	opts = append(opts, http.PathTemplate(pattern))
//	//	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
//	//	if err != nil {
//		return nil, err
//	}
//	return &out, err
//}
//
