package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sanchezta/user-management-lambda/aws"
	"github.com/sanchezta/user-management-lambda/db"
	"github.com/sanchezta/user-management-lambda/models"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	aws.InitAws()

	if !ValidateParameter() {
		fmt.Print("Error en los parámetros. Se debe proporcionar 'SecretName'")

		err := errors.New("parámetros inválidos: se debe proporcionar 'SecretName'")

		return event, err
	}
  
	var data models.SignUp

	for row, att := range event.Request.UserAttributes{


		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email ="+ data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub"+ data.UserUUID)
		}
		
		
	}

	err := db.ReadScret()

	if err != nil {
		fmt.Println("Error al leer el Secret"+err.Error())
		return event, err
	}
	return event, nil
}

// ValidateParameter checks if the required environment variable exists
func ValidateParameter() bool {
	_, exists := os.LookupEnv("SecretName")
	return exists
}
