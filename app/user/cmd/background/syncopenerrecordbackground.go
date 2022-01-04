package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(_ context.Context, _ events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	//for i := 0; i < 60; i++ {
	//	log.Println(i)
	//	time.Sleep(1 * time.Second)
	//}
	return nil, nil
}

func main() {
	lambda.Start(handler)
}
