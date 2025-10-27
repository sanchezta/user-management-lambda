package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanchezta/user-management-lambda/models"
	"github.com/sanchezta/user-management-lambda/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Starting user registration")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	statement := fmt.Sprintf(
		"INSERT INTO users (user_email, user_UUId, user_Dateadd) VALUES ('%s', '%s', '%s')",
		sig.UserEmail,
		sig.UserUUID,
		tools.DateMySQL(),
	)
	fmt.Println(statement)

	_, err = Db.Exec(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
