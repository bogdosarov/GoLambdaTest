package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}

	return events.APIGatewayProxyResponse{
		Body:       "done",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}
