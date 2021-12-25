package main

import (
	env_config "backend/lib/config"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

var echoLambda *echoadapter.EchoLambda

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	abc, a := echoLambda.ProxyWithContext(ctx, req)
	return abc, a
}

func main() {
	env_config.BuildConfig()
	env_config.InitFaunadbClient()

	e := echo.New()
	g := e.Group("/.netlify/functions/user")
	g.POST("/login", login)
	g.GET("/profile", getProfile)
	echoLambda = echoadapter.New(e)

	lambda.Start(handler)
}
