package main

import (
	"BogProject/grpc/pkg/api"
	"BogProject/parcer"
	"BogProject/sqlQuery"
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
)

var NameOfAllCurrency = []string{"Российский рубль", "Австралийский доллар", "Азербайджанский манат", "Армянских драмов", "Белорусский рубль", "Болгарский лев", "Бразильский реал", "Венгерских форинтов", "Вон Республики Корея", "Вьетнамских донгов", "Гонконгский доллар", "Грузинский лари", "Датская крона", "Дирхам ОАЭ", "Доллар США", "Евро", "Египетских фунтов", "Индийских рупий", "Индонезийских рупий", "Казахстанских тенге", "Канадский доллар", "Катарский риал", "Киргизских сомов", "Китайский юань", "Молдавских леев", "Новозеландский доллар", "Новый туркменский манат", "Норвежских крон", "Польский злотый", "Румынский лей", "СДР (специальные права заимствования)", "Сербских динаров", "Сингапурский доллар", "Таджикских сомони", "Таиландских батов", "Турецких лир", "Узбекских сумов", "Украинских гривен", "Фунт стерлингов Соединенного королевства", "Чешских крон", "Шведских крон", "Швейцарский франк", "Южноафриканских рэндов", "Японских иен"}
var NameofAllLitterCode = []string{"RUB", "AUD", "AZN", "AMD", "BYN", "BGN", "BRL", "HUF", "KRW", "VND", "HKD", "GEL", "DKK", "AED", "USD", "EUR", "EGP", "INR", "IDR", "KZT", "CAD", "QAR", "KGS", "CNY", "MDL", "NZD", "TMT", "NOK", "PLN", "RON", "XDR", "RSD", "SGD", "TJS", "THB", "TRY", "UZS", "UAH", "GBP", "CZK", "SEK", "CHF", "ZAR", "JPY"}

var actionButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Узнать курс валют"),
		tgbotapi.NewKeyboardButton("Дешевые авиабилеты"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Аккаунт"),
	),
)

var permissionButton = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("✅", "Подтверждение"),
		tgbotapi.NewInlineKeyboardButtonData("❌", "Отмена"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Пользовательское соглашение", "https://vk.com/feed"),
	),
)

var ChangeInfoAccount = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сменить ФИО", "ФИО"),
		tgbotapi.NewInlineKeyboardButtonData("Сменить Email", "Email"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Сменить Паспорт", "Паспорт"),
		tgbotapi.NewInlineKeyboardButtonData("Сменить Валюту", "Валюта"),
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
			if update.Message.Text != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				// запрос в бд
				sqlQuery.CheckUser(update.Message.From.ID, update.Message.From.UserName)

				if sqlQuery.CheckAgreement(update.Message.From.ID) == false {
					msg.Text = "Подтвердите, что вы прочитали пользовательское соглашение."
					msg.ReplyMarkup = permissionButton
					bot.Send(msg)
				} else if sqlQuery.CheckBlocked(update.Message.From.ID) == false {
					msg.Text = "Вы заблокированы."
					bot.Send(msg)
				} else {
					switch update.Message.Text {
					case "/start":
						msg.Text = "Выберите действие."
						msg.ReplyMarkup = actionButton
						bot.Send(msg)
					case "close":
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						bot.Send(msg)
					case "Узнать курс валют":
						messageCurrency := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите полное название валюты на которую хоите поменять.\nИли найдите через бота @GolangDisterBot валюта.")
						bot.Send(messageCurrency)
						waitNameCurrency(bot, update.Message.Chat.ID, updates, update.Message.From.ID)
					case "Аккаунт":
						messageAccount := tgbotapi.NewMessage(update.Message.Chat.ID, "Здесь будет отображена вся ваша информация об аккаунте")
						email, passport, currency, FIO := sqlQuery.GetAccountInfo(update.Message.From.ID)
						messageAccount.Text += "\nФИО: " + FIO
						messageAccount.Text += "\nEmail: " + email
						messageAccount.Text += "\nПаспорт: " + passport
						messageAccount.Text += "\nВалюта: " + currency
						messageAccount.Text += "\nЕсли у вас есть хотя бы один прочерк, то купить билет вы не сможете."
						messageAccount.ReplyMarkup = ChangeInfoAccount
						bot.Send(messageAccount)
					default:
						msg.Text = "Увы, у нас нет такой команды. Нажмите /start"
						bot.Send(msg)
					}

					logGrpcTelegram(update.Message.Text, update.Message.From.UserName)
				}
			}
		} else if update.CallbackQuery != nil {
			// Получаем данные CallbackQuery
			data := update.CallbackQuery.Data
			// Получаем ID пользователя
			userID := update.CallbackQuery.From.ID

			// Выполняем необходимые действия в зависимости от данных кнопки
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			msg.ReplyMarkup = actionButton

			if data == "Подтверждение" {
				sqlQuery.ChangeAgreemant(userID)
				newmsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Теперь вы можете пользоваться ботом.\nНажмите /start чтобы начать пользоваться.")
				bot.Send(newmsg)
			} else if data == "Отмена" {
				newmsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Вы не можете пользоваться ботом")
				bot.Send(newmsg)
			} else if data == "ФИО" {
				changeFIO(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From.ID, bot, updates)
				bot.Send(msg)
			} else if data == "Email" {
				changeEmail(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From.ID, bot, updates)
				bot.Send(msg)
			} else if data == "Паспорт" {
				changePassport(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From.ID, bot, updates)
				bot.Send(msg)
			} else if data == "Валюта" {
				changeCurrency(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.From.ID, bot, updates)
				bot.Send(msg)
			}
		}
	}
}

func changeFIO(chatID, userID int64, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	message := tgbotapi.NewMessage(chatID, "Введите ваше ФИО.")
	bot.Send(message)

	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text
			sqlQuery.UpdateUserFIO(userID, userInput)

			message.Text = "Ваше ФИО изменено."
			bot.Send(message)

			return
		}
	}

}

func changePassport(chatID, userID int64, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	message := tgbotapi.NewMessage(chatID, "Введите ваше псспорт.")
	bot.Send(message)

	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text
			sqlQuery.UpdateUserPassport(userID, userInput)

			message.Text = "Ваше паспорт изменен."
			bot.Send(message)

			return
		}
	}
}

func changeEmail(chatID, userID int64, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	message := tgbotapi.NewMessage(chatID, "Введите вашу почту.")
	bot.Send(message)

	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text
			sqlQuery.UpdateUserEmail(userID, userInput)

			message.Text = "Ваше email изменен."
			bot.Send(message)

			return
		}
	}
}

func changeCurrency(chatID, userID int64, bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	message := tgbotapi.NewMessage(chatID, "Найдите валюту вашей страны через бота @GolangDisterBot валюта.")
	bot.Send(message)

	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text

			allCurr := parcer.ParcerCurrency()
			var newCurr string
			ok := false

			for _, value := range allCurr {
				if value.Name == userInput {
					newCurr = value.LitterCode
					ok = true
				}
			}

			if ok == true {
				sqlQuery.UpdateUserCurrency(userID, newCurr)
				message.Text = "Ваша основная валюта изменена."
				bot.Send(message)
			} else {
				message.Text = "Такой валюты не существует."
				bot.Send(message)
			}

			return
		} else if update.InlineQuery != nil {
			query := strings.ToLower(update.InlineQuery.Query)
			if query == "валюта" {

				results := make([]interface{}, len(NameOfAllCurrency))

				for i, currency := range NameOfAllCurrency {
					ID, _ := strconv.Atoi(update.InlineQuery.ID)
					article := tgbotapi.NewInlineQueryResultArticle(strconv.Itoa(ID+i), currency, currency)
					article.Description = NameofAllLitterCode[i]
					results[i] = article
				}

				inlineConf := tgbotapi.InlineConfig{
					InlineQueryID: update.InlineQuery.ID,
					IsPersonal:    true,
					CacheTime:     0,
					Results:       results,
				}

				if _, err := bot.Request(inlineConf); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func waitNameCurrency(bot *tgbotapi.BotAPI, chatID int64, updates tgbotapi.UpdatesChannel, userID int64) {

	for update := range updates {
		if update.Message != nil {
			userInput := update.Message.Text
			handleInput(bot, chatID, userInput, updates, userID)
			return
		} else if update.InlineQuery != nil {
			query := strings.ToLower(update.InlineQuery.Query)
			if query == "валюта" {

				results := make([]interface{}, len(NameOfAllCurrency))

				for i, currency := range NameOfAllCurrency {
					ID, _ := strconv.Atoi(update.InlineQuery.ID)
					article := tgbotapi.NewInlineQueryResultArticle(strconv.Itoa(ID+i), currency, currency)
					article.Description = NameofAllLitterCode[i]
					results[i] = article
				}

				inlineConf := tgbotapi.InlineConfig{
					InlineQueryID: update.InlineQuery.ID,
					IsPersonal:    true,
					CacheTime:     0,
					Results:       results,
				}

				if _, err := bot.Request(inlineConf); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func handleInput(bot *tgbotapi.BotAPI, chatID int64, userNameCurr string, updates tgbotapi.UpdatesChannel, userID int64) {

	message := tgbotapi.NewMessage(chatID, "Введите количество валюты.")
	bot.Send(message)

	var userAmountConvert int
	var err error

	for update := range updates {
		if update.Message != nil {
			userAmountConvert, err = strconv.Atoi(update.Message.Text)
			break
		}
	}

	allCurrency := parcer.ParcerCurrency()

	var firstCurrency parcer.Currency
	var secondCurrency parcer.Currency

	fmt.Println(userID)

	currencyFrom := sqlQuery.UserCurrency(userID)

	if err != nil {
		msg := tgbotapi.NewMessage(chatID, "Вы ввели некоректно данные")
		bot.Send(msg)
	} else {
		for _, value := range allCurrency {
			if currencyFrom == value.LitterCode {
				firstCurrency = value
			}
			if userNameCurr == value.Name {
				secondCurrency = value
			}
		}

		fmt.Println(firstCurrency, secondCurrency)

		answer := ((float32(firstCurrency.Rate) * float32(userAmountConvert)) / float32(firstCurrency.Units)) / (float32(secondCurrency.Rate) / float32(secondCurrency.Units))

		// Для примера выводим введенные данные обратно пользователю
		responseMsg := fmt.Sprintf("%s - %0.2f\n%s - %0.2f", firstCurrency.Name, float32(userAmountConvert), secondCurrency.Name, answer)
		msg := tgbotapi.NewMessage(chatID, responseMsg)
		msg.ReplyMarkup = actionButton
		bot.Send(msg)
	}
	return
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
