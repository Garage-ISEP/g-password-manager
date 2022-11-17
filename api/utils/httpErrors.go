package utils

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func (e HttpError) Error() string {
	return e.Message
}
func (e HttpError) ToResponse() *events.APIGatewayProxyResponse {
	body, err := json.Marshal(e)
	if err != nil {
		fmt.Println("Error marshalling error response: ", err.Error())
		return &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       `{"message": "Internal Server Error", "status": 500}`,
		}
	}
	return &events.APIGatewayProxyResponse{
		StatusCode: e.StatusCode,
		Body:       string(body),
	}
}

// Error constructors
func NewBadRequestError(err error) HttpError {
	return HttpError{
		Message:    err.Error(),
		StatusCode: 400,
	}
}
func NewInternalServerError(err error) HttpError {
	fmt.Println("Internal Server Error: ", err.Error())
	return HttpError{
		Message:    `{"message": "Internal Server Error", "status": 500}`,
		StatusCode: 500,
	}
}
