package main

import (
	"log"

	"todo-api/cmd/api"
	"todo-api/db"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Addr:                 "host.docker.internal:3306",
		DBName:               "todos",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage := db.NewMySQLStorage(cfg)

	database, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	storage := db.NewStorage(database)
	api := api.NewAPIServer(":8080", storage)

	api.Start()
}
