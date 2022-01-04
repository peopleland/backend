package http

import (
	ctx "context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

type Group struct {
	*echo.Group
}

type Route struct {
	*echo.Route
}

type HandlerFunc func(Context) error

func NewServer() *Server {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	return &Server{
		e,
	}
}

func InjectAuthorizationMiddleware(h HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		jwtStr := context.Request().Header.Get("authorization")
		return h(Context{
			Context: context,
			Ctx:     ctx.WithValue(ctx.Background(), "authorization", jwtStr),
		})
	}
}

func (s *Server) GroupX(prefix string, m ...echo.MiddlewareFunc) *Group {
	return &Group{s.Echo.Group(prefix, m...)}
}

func (s *Server) Adapter() *echoadapter.EchoLambda {
	return echoadapter.New(s.Echo)
}

func (s *Server) LambdaStart() {
	lambda.Start(func(c ctx.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return s.Adapter().ProxyWithContext(c, req)
	})
}

func (s *Server) HttpStart(address string) {
	for _, route := range s.Echo.Routes() {
		log.Println(route.Name, route.Method, route.Path)
	}

	err := s.Start(address)
	if err != nil {
		return
	}
}

func (g *Group) POSTX(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *Route {
	return &Route{g.Group.POST(path, InjectAuthorizationMiddleware(h), m...)}
}

func (g *Group) GETX(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *Route {
	return &Route{g.Group.GET(path, InjectAuthorizationMiddleware(h), m...)}
}

func (g *Group) PUTX(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *Route {
	return &Route{g.Group.PUT(path, InjectAuthorizationMiddleware(h), m...)}
}

func (g *Group) DELETEX(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *Route {
	return &Route{g.Group.DELETE(path, InjectAuthorizationMiddleware(h), m...)}
}
