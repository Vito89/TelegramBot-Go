package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var host = os.Getenv("HOST")
var port = os.Getenv("PORT")
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var sslmode = os.Getenv("SSLMODE")

var dbInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

func collectData(username string, chatid int64, message string, answer []string) bool {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return false
	}
	defer db.Close()

	//Converting slice with answer to string
	answ := strings.Join(answer, ", ")

	//Creating SQL command
	data := `INSERT INTO users(username, chat_id, message, answer) VALUES($1, $2, $3, $4);`

	//Execute SQL command in database
	if _, err = db.Exec(data, `@`+username, chatid, message, answ); err != nil {
		return false
	}

	return true
}

func createTable() bool {

	//Connecting to database
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return false
	}
	defer db.Close()

	//Creating users Table
	if _, err = db.Exec(`CREATE TABLE users(ID SERIAL PRIMARY KEY, TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT, MESSAGE TEXT, ANSWER TEXT);`); err != nil {
		return false
	}

	return true
}
