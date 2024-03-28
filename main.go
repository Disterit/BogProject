package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

// Структура для хранения информации о валютах
type Currency struct {
	Code   string
	Name   string
	Symbol string
}

// Функция для создания InlineKeyboard с кнопками для выбора валюты
func createCurrencyKeyboard() *tgbotapi.InlineKeyboardMarkup {
	currencies := []Currency{
		{Code: "RUB", Name: "Российский рубль", Symbol: "₽"},
		{Code: "AUD", Name: "Австралийский доллар", Symbol: "A$"},
		{Code: "AZN", Name: "Азербайджанский манат", Symbol: "₼"},
		{Code: "AMD", Name: "Армянских драмов", Symbol: "֏"},
		{Code: "BYN", Name: "Белорусский рубль", Symbol: "Br"},
		{Code: "BGN", Name: "Болгарский лев", Symbol: "лв"},
		{Code: "BRL", Name: "Бразильский реал", Symbol: "R$"},
		{Code: "HUF", Name: "Венгерских форинтов", Symbol: "Ft"},
		{Code: "KRW", Name: "Вон Республики Корея", Symbol: "₩"},
		{Code: "VND", Name: "Вьетнамских донгов", Symbol: "₫"},
		{Code: "HKD", Name: "Гонконгский доллар", Symbol: "HK$"},
		{Code: "GEL", Name: "Грузинский лари", Symbol: "₾"},
		{Code: "DKK", Name: "Датская крона", Symbol: "kr"},
		{Code: "AED", Name: "Дирхам ОАЭ", Symbol: "د.إ"},
		{Code: "USD", Name: "Доллар США", Symbol: "$"},
		{Code: "EUR", Name: "Евро", Symbol: "€"},
		{Code: "EGP", Name: "Египетских фунтов", Symbol: "ج.م"},
		{Code: "INR", Name: "Индийских рупий", Symbol: "₹"},
		{Code: "IDR", Name: "Индонезийских рупий", Symbol: "Rp"},
		{Code: "KZT", Name: "Казахстанских тенге", Symbol: "₸"},
		{Code: "CAD", Name: "Канадский доллар", Symbol: "CA$"},
		{Code: "QAR", Name: "Катарский риал", Symbol: "ر.ق"},
		{Code: "KGS", Name: "Киргизских сомов", Symbol: "сом"},
		{Code: "CNY", Name: "Китайский юань", Symbol: "¥"},
		{Code: "MDL", Name: "Молдавских леев", Symbol: "MDL"},
		{Code: "NZD", Name: "Новозеландский доллар", Symbol: "NZ$"},
		{Code: "TMT", Name: "Новый туркменский манат", Symbol: "TMT"},
		{Code: "NOK", Name: "Норвежских крон", Symbol: "kr"},
		{Code: "PLN", Name: "Польский злотый", Symbol: "zł"},
		{Code: "RON", Name: "Румынский лей", Symbol: "lei"},
		{Code: "XDR", Name: "СДР (специальные права заимствования)", Symbol: "XDR"},
		{Code: "RSD", Name: "Сербских динаров", Symbol: "дин."},
		{Code: "SGD", Name: "Сингапурский доллар", Symbol: "S$"},
		{Code: "TJS", Name: "Таджикских сомони", Symbol: "TJS"},
		{Code: "THB", Name: "Таиландских батов", Symbol: "฿"},
		{Code: "TRY", Name: "Турецких лир", Symbol: "₺"},
		{Code: "UZS", Name: "Узбекских сумов", Symbol: "soʻm"},
		{Code: "UAH", Name: "Украинских гривен", Symbol: "₴"},
		{Code: "GBP", Name: "Фунт стерлингов Соединенного королевства", Symbol: "£"},
		{Code: "CZK", Name: "Чешских крон", Symbol: "Kč"},
		{Code: "SEK", Name: "Шведских крон", Symbol: "kr"},
		{Code: "CHF", Name: "Швейцарский франк", Symbol: "CHF"},
		{Code: "ZAR", Name: "Южноафриканских рэндов", Symbol: "R"},
		{Code: "JPY", Name: "Японских иен", Symbol: "¥"},
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for _, currency := range currencies {
		btn := tgbotapi.NewInlineKeyboardButtonData(currency.Name, currency.Code)
		row := []tgbotapi.InlineKeyboardButton{btn}
		rows = append(rows, row)
	}

	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: rows,
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI("7072095600:AAFeiodIFQGONm7zTItKGLcml6_xMsg9C2w	")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	// Задаем переменную currencies
	currencies := createCurrencyKeyboard()

	for update := range updates {
		if update.CallbackQuery != nil {
			currencyCode := update.CallbackQuery.Data
			// Найдем выбранную валюту по ее коду
			for _, currencyRow := range currencies.InlineKeyboard {
				for _, button := range currencyRow {
					if *button.CallbackData == currencyCode {
						msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
						msg.Text = "Выбрана валюта: " + button.Text + "\nКод: " + *button.CallbackData
						bot.Send(msg)
						break
					}
				}
			}
		} else if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case "/open":
				msg.Text = "Выберите валюту:"
				msg.ReplyMarkup = currencies
				bot.Send(msg)
			}
		}
	}
}
