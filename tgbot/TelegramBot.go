package main

import (
	"BogProject/grpc/pkg/api"
	"BogProject/parcer"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var NameOfAllCurency = "1.Российский рубль     \n2.Австралийский доллар \n3.Азербайджанский манат\n4.Армянских драмов     \n5.Белорусский рубль    \n6.Болгарский лев       \n7.Бразильский реал     \n8.Венгерских форинтов  \n9.Вон Республики Корея \n10.Вьетнамских донгов  \n11.Гонконгский доллар  \n12.Грузинский лари     \n13.Датская крона       \n14.Дирхам ОАЭ          \n15.Доллар США          \n16.Евро                \n17.Египетских фунтов   \n18.Индийских рупий     \n19.Индонезийских рупий \n20.Казахстанских тенге \n21.Канадский доллар    \n22.Катарский риал      \n23.Киргизских сомов    \n24.Китайский юань      \n25.Молдавских леев\n26.Новозеландский доллар\n27.Новый туркменский манат\n28.Норвежских крон\n29.Польский злотый\n30.Румынский лей\n31.СДР (специальные права заимствования)\n32.Сербских динаров\n33.Сингапурский доллар\n34.Таджикских сомони\n35.Таиландских батов\n36.Турецких лир\n37.Узбекских сумов\n38.Украинских гривен\n39.Фунт стерлингов Соединенного королевства\n40.Чешских крон\n41.Шведских крон\n42.Швейцарский франк\n43.Южноафриканских рэндов\n44.Японских иен\n"

var actionButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Узнать курс валют"),
		tgbotapi.NewKeyboardButton("Дешевые авиабилеты"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			logGrpcTelegram(update.Message.Text, update.Message.From.UserName)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				msg.Text = "Выберите действие."
				msg.ReplyMarkup = actionButton
				bot.Send(msg)
			case "close":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.Send(msg)
			case "Узнать курс валют":
				messageCurrency := tgbotapi.NewMessage(update.Message.Chat.ID, NameOfAllCurency)
				bot.Send(messageCurrency)
				waitForInput(bot, update.Message.Chat.ID, updates)
			default:
				msg.Text = "Увы у нас нет такой команды нажмите /start"
				bot.Send(msg)
			}
		}
	}
}

func waitForInput(bot *tgbotapi.BotAPI, chatID int64, updates tgbotapi.UpdatesChannel) {

	message := tgbotapi.NewMessage(chatID, "Введите номер валюты, количество и номер валюты, на которую хотите поменять.")
	bot.Send(message)

	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text

			handleInput(bot, chatID, userInput)

			return
		}
	}
}

func handleInput(bot *tgbotapi.BotAPI, chatID int64, userInput string) {

	allCurrency := parcer.ParcerCurrency()

	var firstCurrency parcer.Currency
	var secondCurrency parcer.Currency

	parts := strings.Split(userInput, " ")

	var currencyFrom int
	var amount int
	var currencyTo int
	var err error

	if len(parts) == 3 {
		currencyFrom, err = strconv.Atoi(parts[0])
		amount, err = strconv.Atoi(parts[1])
		currencyTo, err = strconv.Atoi(parts[2])
	}

	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Вы ввели некоректно данные")
		bot.Send(msg)
	} else {
		for i, value := range allCurrency {
			if currencyFrom == i+1 {
				firstCurrency = value
			}
			if currencyTo == i+1 {
				secondCurrency = value
			}
		}

		answer := ((float32(firstCurrency.Rate) * float32(amount)) / float32(firstCurrency.Units)) / (float32(secondCurrency.Rate) / float32(secondCurrency.Units))

		// Для примера выводим введенные данные обратно пользователю
		responseMsg := fmt.Sprintf("%s - %0.2f\n%s - %0.2f", firstCurrency.Name, float32(amount), secondCurrency.Name, answer)
		msg := tgbotapi.NewMessage(chatID, responseMsg)
		msg.ReplyMarkup = actionButton
		bot.Send(msg)
	}
}

func logGrpcTelegram(messageUser, username string) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось установить соединение: %v", err)
	}
	defer conn.Close()

	client := api.NewTelegramBotClient(conn)

	request := &api.MessageRequestTelegram{
		User:    username,
		Message: messageUser,
	}

	response, err := client.GetMessages(context.Background(), request)
	if err != nil {
		log.Fatalf("Ошибка при вызове метода GetMessages: %v", err)
	}

	// Выводим полученные сообщения на консоль.
	for _, message := range response.Messages {
		fmt.Println(message.User, message.Message, "aboba")
	}

}
