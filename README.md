# TelegramBot-Go

 This is a Telegram bot which receives information from user which he want found in Wikipedia and return information which he found in Wikipedia.

![example work of bot](https://github.com/trigun117/TelegramBot-Go/blob/master/example.jpg)
# Getting Started
### Telegram
If you want to use /number_of_users command your need to connect bot to database and add this command to your bot using BotFather: **number_of_users - number of users who used bot**

/number_of_users command sending message to user which contain number of users who used bot

### Docker Hub
https://hub.docker.com/r/trigun117/wikipedia-telegram-bot/
### Docker
#### Language
You can specify language
```
-e LANGUAGE=search_language
Examples: 
Russian: docker run -e LANGUAGE=ru -e TOKEN=set_your_bot_token -d trigun117/wikipedia-telegram-bot
France: docker run -e LANGUAGE=fr -e TOKEN=set_your_bot_token -d trigun117/wikipedia-telegram-bot
```
Default is English

#### Image
If you dont want to use databse run:
```
docker run -e TOKEN=set_your_bot_token -d trigun117/wikipedia-telegram-bot
```
If you want to use database run:
```
docker run \
-e DB_SWITCH=on \
-e TOKEN=set_your_bot_token \
-e HOST=set_your_database_host \
-e PORT=set_your_database_port \
-e USER=set_your_database_user \
-e PASSWORD=set_your_database_password \
-e DBNAME=set_your_database_name \
-e SSLMODE=set_your_database_sslmode(disable or enable, default disable) \
-d trigun117/wikipedia-telegram-bot
```
### Docker-Compose
Set environment variables in docker-compose.yml
```
services:

  db:
    image: postgres
    environment:
      POSTGRES_PASSWORD: test

  bot:
    image: trigun117/wikipedia-telegram-bot
    environment:
      DB_SWITCH: "on"
      TOKEN: set_your_bot_token
      HOST: db
      PORT: 5432
      USER: postgres
      PASSWORD: test
      DBNAME: postgres
      SSLMODE: disable
```
Than move to /docker-compose folder and run:
```
docker-compose up --build
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
