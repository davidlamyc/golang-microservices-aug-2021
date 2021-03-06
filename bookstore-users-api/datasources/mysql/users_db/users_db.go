package users_db

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		"user",
		"password",
		"127.0.0.1:3306",
		"users",
	)

	// datasourceName := fmt.Sprintf(
	// 	"%s:%s@tcp(%s)/%s?charset=utf8",
	// 	username,
	// 	password,
	// 	host,
	// 	schema,
	// )

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("database successfully configured.")
}