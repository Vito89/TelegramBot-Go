package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"

	"github.com/Syfaro/telegram-bot-api"
)

//SearchResults structure with query and all Result data
type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}

//Result structure which contains parsed results
type Result struct {
	Name, Description, URL string
}

//UnmarshalJSON set data to strcuture
func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}

//urlEncoded encodes a string to be used in a query part of a URL
func urlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func wikipediaAPI(request string) (answer []string) {

	s := make([]string, 3)

	//Sending request
	if response, err := http.Get(request); err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()

		//Reading answer
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		//Unmarshal answer and set it to SearchResults struct
		sr := &SearchResults{}
		if err = json.Unmarshal([]byte(contents), sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
		}

		//Check if struct is not empty
		if !sr.ready {
			s[0] = "Something going wrong, try to change your question"
		}

		//Loop through Results struct and assigning data to s slice
		for i := range sr.Results {
			s[i] = sr.Results[i].URL
		}
	}
	return s
}

func telegramBot() {

	//Create bot
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
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
