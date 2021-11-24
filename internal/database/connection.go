package database

import (
	"fmt"
  _ "github.com/lib/pq"
          "database/sql"
	"errors"
)

var Conn *sql.DB

func ConnectionInfo() string {
	host, port, user, password, database := Credentials()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, database)
  return psqlInfo
}

func SetConnection() error{
	db, err := sql.Open("postgres", ConnectionInfo())
	if err != nil {
		return errors.New("401")
	}
	Conn = db
	return errors.New("200")
}

func CloseConnection(){
	defer Conn.Close()
}

