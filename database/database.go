/*
Author: Filip Johansson
Email:  filip.carl.johansson@gmail.com
Github: fippli

Description:
Database connection to postgresql database.
Configurations for the database in ./db.config.json.

*/

package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"../util"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

var (
	DB *sql.DB
)

func Connection() *sql.DB {

	var config DBConfig
	configFile := util.ReadFile("database/db.config.json")
	err := json.Unmarshal([]byte(configFile), &config)

	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
