FROM golang

COPY ./  /app

WORKDIR /app

RUN apt update && apt install -y cron syslog-ng

RUN go get

RUN go build
