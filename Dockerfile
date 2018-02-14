FROM alpine
COPY TelegramBot-Go .
RUN apk add --no-cache ca-certificates &&\
    chmod +x TelegramBot-Go
EXPOSE 80/tcp
CMD [ "./TelegramBot-Go" ]