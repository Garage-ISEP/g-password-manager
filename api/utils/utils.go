package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Handler func(ctx context.Context, request events.APIGatewayProxyRequest) (interface{}, error)

func HandleAWSProxy(handler Handler, ctx context.Context, request events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {

	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error", r)
		}
	}()

	res, err := handler(ctx, request)
	if err != nil {
		if httpError, ok := err.(HttpError); ok {
			return httpError.ToResponse()
		} else {
			return NewInternalServerError(err).ToResponse()
		}
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

func ValidateBody(body string, object interface{}) error {
	if err := json.Unmarshal([]byte(body), object); err != nil {
		fmt.Printf("Error unmarshalling body: %v", err)
		return NewInternalServerError(fmt.Errorf("Error unmarshalling body"))
	}
	// Iterate over all object fields to get tags
	t := reflect.TypeOf(object).Elem()
	v := reflect.ValueOf(object).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Check if field is required
		if field.Tag.Get("required") == "true" {
			// Check if field is empty
			if v.Field(i).IsZero() {
				return NewBadRequestError(fmt.Errorf("field %s is required", field.Tag.Get("json")))
			}
		}
	}
	return nil
}

func LambdaStart(handler Handler) {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		res := HandleAWSProxy(handler, ctx, request)
		return res, nil
	})
}
