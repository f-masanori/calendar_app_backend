package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var ConnectionError error

type SqlHandler struct {
	DB *sql.DB
}

func NewSqlHandler(datasource string) *SqlHandler {

	Db, ConnectionError = sql.Open("mysql", datasource)
	fmt.Println("DB-information")
	fmt.Println("datasource: " + datasource)
	fmt.Println("______________")
	if ConnectionError != nil {
		log.Fatal("error connecting to database: ", ConnectionError)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.DB = Db
	return sqlHandler
}

func CloseConn() {
	Db.Close()
}
