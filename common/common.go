package common

import (
	"os"

	"github.com/sanchezta/user-management-lambda/models"
	"github.com/sanchezta/user-management-lambda/secrets"
)


var SecretModel models.SecretRDSJson
var err error

func ReadScret () error{

	SecretModel, err = secrets.GetSecret(os.Getenv("SecretName"))

	return  err
}
