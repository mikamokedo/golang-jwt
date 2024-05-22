package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)


func NewSqlDb(config *mysql.Config) (*sql.DB,error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	return db, err
}