package db

import (
	"database/sql"
	"fmt"
	"log"
)

type DBConfig struct {
	Host string
	Port uint
	User string
	Password string
	Database string
}

func GetConnection(configs DBConfig) *sql.DB {
	log.Printf("Creating the DB connection with %s\n", configs.Host)
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s",
    configs.Host, configs.Port, configs.User, configs.Password, configs.Database)

	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("Failed to connect with database")
		panic(err)
	}
	log.Println("Connection established successfully")
	return dbConnection
}

