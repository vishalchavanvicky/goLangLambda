package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Product struct {
	Name     *string `json:"name"`
	Quantity *string `json:"quantity"`
}

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var product Product
	err := json.Unmarshal([]byte(request.Body), &product)

	if err != nil {
		log.Print(err.Error())
		return events.APIGatewayProxyResponse{}, err
	}

	msg := fmt.Sprintf("%v - %v", *product.Name, *product.Quantity)
	log.Print(msg)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       msg,
	}, nil
}
