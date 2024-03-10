package main

import (
	"context"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	awsgo "github.com/juparefe/Golang-Ecommerce-User/awsgo"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(context context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.StartAWS()
	return
}
