package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ajiepangestu/go-rest-api-example/config"
)

var DB *sql.DB

func Connect() error {
	var err error
	// DSN format = username:password@protocol(address)/dbname?param=value
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_PROTOCOL"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"),
	)
	DB, err = sql.Open("mysql", dsn)
	if err = DB.Ping(); err != nil {
		return err
	}
	InitDatabase()
	fmt.Println("Connection Opened to Database")
	return nil
}
