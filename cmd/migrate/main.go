package main

import (
	"fmt"
	"jwt-api/config"
	"jwt-api/db"
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)


func main(){
	config := config.ENV;
	db,err := db.NewSqlDb(&mysqlCfg.Config{
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}
    m, err := migrate.NewWithDatabaseInstance(
        "file://cmd/migrate/migrations",
        "mysql", 
        driver,
    )

	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}


}