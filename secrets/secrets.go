package secrets

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/sanchezta/user-management-lambda/models"

	awsgo "github.com/sanchezta/user-management-lambda/aws"
)



func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	
	fmt.Printf(" > Solicitando secreto: %s\n", nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	if err != nil {
		fmt.Printf(" > Error al obtener secreto: %s - %v\n", nameSecret, err)
		return dataSecret, fmt.Errorf("error al obtener el secreto %s: %w", nameSecret, err)
	}

	if key.SecretString == nil {
		return dataSecret, fmt.Errorf("el secreto %s no tiene valor de cadena", nameSecret)
	}

	if err := json.Unmarshal([]byte(*key.SecretString), &dataSecret); err != nil {
		fmt.Printf(" > Error al deserializar secreto: %s - %v\n", nameSecret, err)
		return dataSecret, fmt.Errorf("error al deserializar el secreto %s: %w", nameSecret, err)
	}

	fmt.Printf(" > Lectura de secreto exitosa: %s\n", nameSecret)
	return dataSecret, nil
}
