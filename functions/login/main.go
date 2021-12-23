package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LoginPayload struct {
	Address       string `json:"address"`
	Signature     string `json:"signature"`
	OriginMessage string `json:"origin_message"`
}

type LoginResponseBody struct {
	Jwt string `json:"jwt"`
}

func build500Response(errorMsg string) (*events.APIGatewayProxyResponse, error) {
	var jsonBody []byte
	var err error
	boby := make(map[string]string)
	boby["error"] = errorMsg
	jsonBody, err = json.Marshal(boby)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      500,
		Headers:         map[string]string{"Content-Type": "text/json"},
		Body:            string(jsonBody),
		IsBase64Encoded: false,
	}, nil
}

func buildJsonResponse(data interface{}) (*events.APIGatewayProxyResponse, error) {
	var jsonBody []byte
	var err error

	boby := make(map[string]interface{})
	boby["data"] = data

	jsonBody, err = json.Marshal(boby)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "text/json"},
		Body:            string(jsonBody),
		IsBase64Encoded: false,
	}, nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		return build500Response("request.http_method.error")
	}

	var loginPayload LoginPayload
	var err error
	if err = json.Unmarshal([]byte(request.Body), &loginPayload); err != nil {
		return build500Response("request.format.error")
	}

	return buildJsonResponse(LoginResponseBody{Jwt: "hello"})
}

func main() {
	lambda.Start(handler)
}
