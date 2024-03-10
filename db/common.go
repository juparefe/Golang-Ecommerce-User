package db

import (
	"os"

	"github.com/juparefe/Golang-Ecommerce-User/models"
	"github.com/juparefe/Golang-Ecommerce-User/secretmngr"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secretmngr.GetSecret(os.Getenv("SecretName"))
	return err
}
