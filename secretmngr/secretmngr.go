package secretmngr

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	awsgo "github.com/juparefe/Golang-Ecommerce-User/awsgo"
	models "github.com/juparefe/Golang-Ecommerce-User/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	secretData := models.SecretRDSJson{}
	fmt.Println("Getting secret: " + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Config)
	secretValue, err := svc.GetSecretValue(awsgo.Context, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println("Error getting secret value:", err.Error())
		return secretData, err
	}
	// Procesar el secretValue y guardarlo en secretData
	err = json.Unmarshal([]byte(*secretValue.SecretString), &secretData)
	if err != nil {
		fmt.Println("Error unmarshalling secret: ", secretName, err.Error())
		return secretData, err
	}
	fmt.Println("Secret data OK: ", secretData)
	return secretData, nil
}
