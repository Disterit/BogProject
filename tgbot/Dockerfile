FROM golang:1.22

WORKDIR /tgbot

COPY go.mod /tgbot
COPY go.sum /tgbot
COPY /tgbot /tgbot

RUN go mod download

CMD ["go", "run", "TelegramBot.go"]