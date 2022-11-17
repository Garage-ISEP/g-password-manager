package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"
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
	fmt.Println("Handler executed")
	if err != nil {
		if httpError, ok := err.(HttpError); ok {
			return httpError.ToResponse()
		} else {
			return NewInternalServerError(err).ToResponse()
		}
	}

	var data []byte
	if res != nil {
		data, err = json.Marshal(res)
		if err != nil {
			return NewInternalServerError(err).ToResponse()
		}
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(data),
	}
}

func ValidateBody(body string, object interface{}) error {
	if err := json.Unmarshal([]byte(body), object); err != nil {
		return NewBadRequestError(fmt.Errorf("Error unmarshalling body: %s", err.Error()))
	}
	return validateObject(object)
}

// Parse a param map into a struct
func ValidateQueryParams(params *map[string]string, object interface{}) error {
	if err := mapstructure.Decode(params, object); err != nil {
		return NewBadRequestError(fmt.Errorf("Error umarshalling query params: %s", err.Error()))
	}
	return validateObject(object)
}

func validateObject(object interface{}) error {
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
