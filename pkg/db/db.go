package db

import (
	"database/sql"
	"fmt"
	"log"
	"swift/internal/configs"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenConnection()(*sql.DB, error){
	var err error
	pgConfig := configs.Settings.Db
	options := fmt.Sprintf("user = %s dbname = %s password = %s sslmade = disable",pgConfig.User, pgConfig.DBname, pgConfig.Password)
	DB, err = sql.Open("postgres", options)
	if err != nil {
		return nil, err
	}
	return DB, nil
} 

func CloseConnection(){
	err := DB.Close()
	if err != nil {
		log.Fatalln("ERROR while closing connection")
	}
}