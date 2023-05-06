package db

import (
	"database/sql"

	"github.com/PBKKE08/FP-BE/echo-rest/config"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func Init(){
	config := config.GetConfig()
	connStr := "user=" + config.DB_USERNAME + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_NAME + " host=" + config.DB_HOST + " port=" + config.DB_PORT + " sslmode=disable"

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetDb() *sql.DB{
	return db
}