package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
)

func collectData(username string, chatid int64, message string, answer []string) bool {
	//CREATE TABLE users(ID SERIAL PRIMARY KEY, TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT, MESSAGE TEXT, ANSWER TEXT);

	//Connecting to database
	db, err := sql.Open("postgres", `host= database host port= databse port user= username password= password dbname= database name sslmode= enable or disable(default disable)`)
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
