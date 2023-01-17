package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/palomabarroso/go-web-application/helpers"
)

func ConnectDB() *sql.DB {
	user := helpers.GoDotEnvVariable("DB_USER")
	pwd := helpers.GoDotEnvVariable("DB_PWD")
	dbName := helpers.GoDotEnvVariable("DB_NAME")
	host := helpers.GoDotEnvVariable("DB_HOST")
	connection := "user=" + user + " dbname=" + dbName + " password=" + pwd + " host=" + host + " sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
