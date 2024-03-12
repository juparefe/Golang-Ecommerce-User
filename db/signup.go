package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juparefe/Golang-Ecommerce-User/models"
	"github.com/juparefe/Golang-Ecommerce-User/tools"
)

// Conectar a la base de datos y ejecutar el script para insertar el nuevo usuario
func SignUpInDb(sig models.SignUp) error {
	fmt.Println("Start registration")
	err := DbConnect()
	if err != nil {
		fmt.Println("Error connecting to Database:", err.Error())
		return err
	}
	defer Db.Close()
	SqlScript := "INSERT INTO users (User_Email,User_UUID,User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DateMySQL() + "')"
	fmt.Println("Script:", SqlScript)
	_, err = Db.Exec(SqlScript)
	if err != nil {
		fmt.Println("Error executing Insert script in Database:", err.Error())
		return err
	}
	fmt.Println("Registration complete")
	return nil
}
