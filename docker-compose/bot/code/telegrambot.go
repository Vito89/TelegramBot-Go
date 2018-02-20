package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"os"
	"reflect"
	"time"
)

func telegramBot() {

	//Create bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	//Set update timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Get updates from bot
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		//Check if message from user is text
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			//Set search language
			language := os.Getenv("LANGUAGE")

			//Create search url
			ms, _ := urlEncoded(update.Message.Text)

			url := ms
			request := "https://" + language + ".wikipedia.org/w/api.php?action=opensearch&search=" + url + "&limit=3&origin=*&format=json"

			//assigning value of answer slice to variable message
			message := wikipediaAPI(request)

			if os.Getenv("DB_SWITCH") == "on" {

				//Putting username, chat_id, message, answer to database
				if err := collectData(update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text, message); err != true {

					//Send message
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error, but bot still working")
					bot.Send(msg)
				}
			}

			//Loop throug message slice
			for _, val := range message {

				//Send message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, val)
				bot.Send(msg)
			}
		} else {

			//Send message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use words for search")
			bot.Send(msg)
		}
	}
}

func main() {

	time.Sleep(15 * time.Second)

	//Creating Table
	if os.Getenv("DB_SWITCH") == "on" {
		if err := createTable(); err != true {
			panic(err)
		}
	}

	//Call Bot
	telegramBot()
}
