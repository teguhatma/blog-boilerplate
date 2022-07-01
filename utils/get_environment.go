package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func DatabaseVariable() (string, string) {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSLMODE")

	return driver, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pwd, dbName, ssl)
}

func ServerVariable() (string, int, string) {
	name := os.Getenv("SERVER_NAME")
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal("Error casting on Server Environment")
	}
	version := os.Getenv("SERVER_VERSION")

	return name, port, version
}
