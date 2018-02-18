# TelegramBot-Go

 This is a Telegram bot which receives information from user which he want found in Wikipedia and return information which he found in Wikipedia.

![example work of bot](https://github.com/trigun117/TelegramBot-Go/blob/master/example.jpg)
# Getting Started

### Docker

Run run:
```
docker run -e TOKEN=set_your_bot_token -e HOST=set_your_database_host -e PORT=set_your_database_port -e USER=set_your_database_user -e PASSWORD=set_your_database_password -e DBNAME=set_your_database_name -e SSLMODE=set_your_database_sslmode(disable or enable, default disable) -d trigun117/wikipedia-telegram-bot
```
### Docker-Compose
Set environment variables
```
services:

  db:
    build: ./database
    environment:
      POSTGRES_PASSWORD: test

  bot:
    image: trigun117/wikipedia-telegram-bot
    environment:
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
docker-compose up
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
