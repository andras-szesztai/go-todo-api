package main

import (
	"log"

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
	sqlStorage := NewMySQLStorage(cfg)

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStorage(db)
	api := NewAPIServer(":8080", store)

	api.Start()
}
