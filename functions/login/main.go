package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
	"os"
	"strings"

	"backend/lib/helper"
	"github.com/aws/aws-lambda-go/events"
	"github.com/spf13/viper"
)

type LoginPayload struct {
	Address       string `json:"address"`
	Signature     string `json:"signature"`
	OriginMessage string `json:"origin_message"`
}

type LoginResponseBody struct {
	Jwt string `json:"jwt"`
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		return helper.Build500Response("request.http_method.error")
	}

	var loginPayload LoginPayload
	var err error
	if err = json.Unmarshal([]byte(request.Body), &loginPayload); err != nil {
		return helper.Build500Response("request.format.error")
	}

	body, err1 := process(&loginPayload)
	if err != nil {
		return helper.Build500Response(err1.Message)
	}

	return helper.BuildJsonResponse(body)
}

func process(loginPayload *LoginPayload) (*LoginResponseBody, *helper.CommonError) {
	verifed := helper.VerifyEip191Sig(loginPayload.Address, loginPayload.OriginMessage, loginPayload.Signature)
	if !verifed {
		return nil, &helper.CommonError{Message: "request.verify.error"}
	}
	address := strings.ToLower(loginPayload.Address)

	claims := jwt.MapClaims{"address": address}
	jwtStr, err := helper.EncodeJwt(claims, viper.GetString("JWT_RSA_PRIVATE_KEY_PEM"), int64(86400))
	if err != nil {
		return nil, &helper.CommonError{Message: "request.jwt.error"}
	}
	return &LoginResponseBody{Jwt: jwtStr}, nil
}

func main() {
	viper.Set("JWT_RSA_PRIVATE_KEY_PEM", os.Getenv("JWT_RSA_PRIVATE_KEY_PEM"))
	lambda.Start(handler)
}
