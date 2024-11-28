package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DBConfig struct {
	Host string
	Port uint
	User string
	Password string
	Database string
}

func GetConnection() *sql.DB {
	configs := getDBConfigurations()
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

func getDBConfigurations () DBConfig {
	var dbPort, _ = strconv.ParseUint(os.Getenv("DB_PORT"), 10, 16)

	return DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: uint(dbPort),
		User: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

