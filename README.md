# TelegramBot-Go

 This is a Telegram bot which receives information from user which he want found in Wikipedia and return information which he found in Wikipedia.

![example work of bot](https://github.com/trigun117/TelegramBot-Go/blob/master/example.jpg)
# Getting Started

Download code

Set your bot token here
```
//Create bot
bot, err := tgbotapi.NewBotAPI("TOKEN")
 ```
Set database credential
```
db, err := sql.Open("postgres", `host= database host port= databse port user= username password= password dbname= database name sslmode= enable or disable(default disable)`)
```
Than
 
```
go build
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
