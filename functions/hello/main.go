package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/giornetta/lambda-go/pkg/helpers"
)

// Handler handles the function
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return helpers.Response(map[string]string{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	}, http.StatusOK)
}

func main() {
	lambda.Start(Handler)
}
