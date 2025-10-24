package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sanchezta/user-management-lambda/aws"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	aws.InitAws()

	if !ValidateParameter() {
		fmt.Print("Error in parameters. 'SecretName' must be provided")

		err := errors.New("invalid parameters: 'SecretName' must be provided")

		return event, err
	}

	return event, nil
}

// ValidateParameter checks if the required environment variable exists
func ValidateParameter() bool {
	_, exists := os.LookupEnv("SecretName")
	return exists
}
