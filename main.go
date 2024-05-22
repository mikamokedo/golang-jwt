package main

import (
	"database/sql"
	"fmt"
	"jwt-api/cmd/api"
	"jwt-api/config"
	"jwt-api/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main(){
	config := config.ENV;
	db,err := db.NewSqlDb(&mysql.Config{
		User:   config.DBUser,
		Passwd: config.DBPasswd,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", config.DBHost, config.DBPort),
		DBName: config.DBName,
		ParseTime: true,
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initDB(db);
	server := api.NewAPIServer(config.Port,db);

	if error := api.RunServer(server); error != nil {
		log.Fatal(error)
	}
}

func initDB(db *sql.DB) {
	if error:= db.Ping(); error != nil {
		log.Fatal(error)
	}
	log.Println("Connected database")

}