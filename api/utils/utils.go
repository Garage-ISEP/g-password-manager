package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (*interface{}, *error)

func HandleAWSProxy(handler Handler, ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in HandleAWSProxy", r)
		}
	}()
	res, err := handler(ctx, request)
	if err != nil {
		return nil, *err
	}
	var jsonData string = ""
	if res != nil {
		data, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		jsonData = string(data)
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonData),
	}, nil
}
