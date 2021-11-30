package database

import (
	"os"
)


func Credentials() (string, int, string, string, string) {
	host := "localhost"
	port := 5432
	user, _ := os.LookupEnv("DATABASE_USER")
	password, _ := os.LookupEnv("DATABASE_USER_PASSWORD")
	database, _ := os.LookupEnv("DATABASE_DATABASE")
	return host, port, user, password, database
}
