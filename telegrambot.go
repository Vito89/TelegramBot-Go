package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"reflect"
)

func telegramBot() {

	//Create bot
	bot, err := tgbotapi.NewBotAPI("TOKEN")
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

			//Create search url
			ms, _ := urlEncoded(update.Message.Text)

			url := ms
			request := "https://en.wikipedia.org/w/api.php?action=opensearch&search=" + url + "&limit=3&origin=*&format=json"

			//assigning value of answer slice to variable message
			message := wikipediaAPI(request)

			//Putting username, chat_id, message, answer to database
			collectData(update.Message.Chat.UserName, update.Message.Chat.ID, ms, message)

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

	//Call Bot
	telegramBot()
}
