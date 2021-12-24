package main

import (
	"backend/lib/config"
	"backend/lib/helper"
	"backend/lib/models"
	"encoding/json"
	"errors"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang-jwt/jwt"
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
	if err1 != nil {
		return helper.Build500Response(err1.Error())
	}

	return helper.BuildJsonResponse(body)
}

func process(loginPayload *LoginPayload) (*LoginResponseBody, error) {
	verifed := helper.VerifyEip191Sig(loginPayload.Address, loginPayload.OriginMessage, loginPayload.Signature)
	if !verifed {
		return nil, errors.New("request.verify.error")
	}
	address := strings.ToLower(loginPayload.Address)

	_, err1 := models.FindOrCreateUser(env_config.FaunadbClient, address)
	if err1 != nil {
		return nil, errors.New("request.db.error")
	}

	claims := jwt.MapClaims{"address": address}
	jwtStr, err2 := helper.EncodeJwt(claims, env_config.Conf.JwtRsaPrivateKeyPem, int64(86400))
	if err2 != nil {
		return nil, errors.New("request.jwt.error")
	}
	return &LoginResponseBody{Jwt: jwtStr}, nil
}

func main() {
	env_config.BuildConfig()
	env_config.InitFaunadbClient()
	lambda.Start(handler)
}
