package helper

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

type CommonError struct {
	Message string
}

func (c CommonError) Error() string {
	return c.Message
}

func (c CommonError) RuntimeError() {
	fmt.Println(c.Message)
}

func Build500Response(errorMsg string) (*events.APIGatewayProxyResponse, error) {
	var jsonBody []byte
	var err error
	body := make(map[string]string)
	body["error"] = errorMsg
	jsonBody, err = json.Marshal(body)
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

func BuildJsonResponse(data interface{}) (*events.APIGatewayProxyResponse, error) {
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
