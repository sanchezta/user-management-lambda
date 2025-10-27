package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanchezta/user-management-lambda/models"
	"github.com/sanchezta/user-management-lambda/secrets"
)


var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadScret () error{

	SecretModel, err = secrets.GetSecret(os.Getenv("SecretName"))
	
	return  err
}


func DbConnect() error {
    Db, err = sql.Open("mysql", ConnStr(SecretModel))
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    err = Db.Ping()
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    fmt.Println("Conexión a MySQL exitosa")
    return nil
}


func ConnStr(keys models.SecretRDSJson) string{
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "gambit"

	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",dbUser,authToken,dbEndpoint,dbName)

	fmt.Println(dns)

	return dns
}
