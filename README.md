# TelegramBot-Go

 This is a Telegram bot which receives information from user which he want found in Wikipedia and return information which he found in Wikipedia.

![example work of bot](https://github.com/trigun117/TelegramBot-Go/blob/master/example.jpg)
# Getting Started

### Creating executable file

Set your bot token in telegrambot.go:
```
//Create bot
bot, err := tgbotapi.NewBotAPI("TOKEN")
 ```
Set database credential in db.go:
```
db, err := sql.Open("postgres", `host= database host port= databse port user= username password= password dbname= database name sslmode= enable or disable(default disable)`)
```
Than move to /docker-compose/bot/code folder and run:
 
```
go build
```
### Docker Image
Set your bot token in telegrambot.go:
```
//Create bot
bot, err := tgbotapi.NewBotAPI("TOKEN")
 ```
Set database credential in db.go:
```
db, err := sql.Open("postgres", `host= database host port= databse port user= username password= password dbname= database name sslmode= enable or disable(default disable)`)
```
Move to /docker-compose/bot/code folder and run:
```
GOOS=linux GOARCH=amd64 go build
```
Than move to /docker-compose/bot folder and run:
```
docker build -t image_name -f Dockerfile .
```
### Docker-Compose
Set your bot token in telegrambot.go:
```
//Create bot
bot, err := tgbotapi.NewBotAPI("TOKEN")
 ```
 Set this credential in db.go:
```
db, err := sql.Open("postgres", `host=db port=5432 user=postgres password=wikitelegrambot dbname=postgres sslmode=disable`)
```
Move to /docker-compose/bot/code folder and run:
```
GOOS=linux GOARCH=amd64 go build
```
Than move to /docker-compose folder and run:
```
docker-compose up
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
