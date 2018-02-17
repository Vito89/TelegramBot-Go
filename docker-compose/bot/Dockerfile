FROM debian
COPY /code/code .
RUN apt-get update &&\
    apt-get upgrade -y &&\
    apt-get dist-upgrade -y &&\
    apt-get install ca-certificates &&\
    chmod +x code
EXPOSE 80/tcp
CMD [ "./code" ]