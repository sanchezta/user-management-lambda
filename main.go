package main

import (
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sanchezta/user-management-lambda/aws"

	"github.com/aws/aws-lambda-go/events"

	"context"
)
func main (){

	lambda.Start(EjecutoLambda)

}


func EjecutoLambda (ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error){

	aws.InitAws()

	return event, nil

}
