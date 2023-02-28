package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler{})
}

type Handler struct{}

func (h Handler) Invoke(ctx context.Context, payload []byte) (op []byte, err error) {
	var res events.APIGatewayV2HTTPResponse
	res.Body = "Hello, World!"
	return json.Marshal(res)
}
