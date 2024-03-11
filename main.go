package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	awsgo "github.com/juparefe/Golang-Ecommerce-User/awsgo"
	"github.com/juparefe/Golang-Ecommerce-User/db"
	models "github.com/juparefe/Golang-Ecommerce-User/models"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(context context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// Obtener context y config de AWS
	awsgo.StartAWS()
	// Validar que este el parametro secretName en las variables de entorno
	if !ValidateParameters() {
		fmt.Println("SecretName parameter is missing")
		err := errors.New("SecretName parameter is missing")
		return event, err
	}
	// Obtener email y UUID del evento
	var data models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email: ", data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("UUID: ", data.UserUUID)
		}
	}
	// Leer el secreto
	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret: ", err.Error())
		return event, err

	}
	err = db.SignUpInDb(data)
	return event, nil
}

func ValidateParameters() bool {
	var bringParameter bool
	_, bringParameter = os.LookupEnv("SecretName")
	return bringParameter
}
