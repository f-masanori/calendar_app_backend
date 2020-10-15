package database

import (
	"database/sql"
	"fmt"
	"log"

	"golang/conf"

	_ "github.com/go-sql-driver/mysql"
)

var TestDb *sql.DB
var TestConnectionError error

func TestNewSqlHandler() *SqlHandler {
	//configからDBの読み取り
	//read DB config-information from config
	connectionCmd := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Dbname,
	)

	Db, ConnectionError = sql.Open(conf.Database.Drivername, connectionCmd)
	if ConnectionError != nil {
		log.Fatal("error connecting to database: ", ConnectionError)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.DB = Db

	return sqlHandler
}
