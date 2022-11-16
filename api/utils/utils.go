package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error)

func internalServerError(err error) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}
}
func HandleAWSProxy(handler Handler, ctx context.Context, request events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {

	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
		}
	}()

	res, err := handler(ctx, request)
	if err != nil {
		return internalServerError(err)
	}

	var data []byte
	if res != nil {
		data, _ = json.Marshal(res)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}
}
