package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

func collectData(username string, chatid int64, message string, answer []string) bool {
	//CREATE TABLE users(ID SERIAL PRIMARY KEY, TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT, MESSAGE TEXT, ANSWER TEXT);
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("SSLMODE")

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
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
