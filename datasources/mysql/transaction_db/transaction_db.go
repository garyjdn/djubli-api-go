package transaction_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_transactions_username = "mysql_transactions_username"
	mysql_transactions_password = "mysql_transactions_password"
	mysql_transactions_host     = "mysql_transactions_host"
	mysql_transactions_schema   = "mysql_transactions_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysql_transactions_username)
	password = os.Getenv(mysql_transactions_password)
	host     = os.Getenv(mysql_transactions_host)
	schema   = os.Getenv(mysql_transactions_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		//panic(err)
	}
	if err := Client.Ping(); err != nil {
		//panic(err)
	}
	log.Println("database successfully configured")
}
