package parcer

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Currency struct {
	digitalCode int
	litterCode  string
	Units       int
	Name        string
	Rate        float32
}

func ParcerCurrency() []Currency {
	// Загружаем страницу
	res, err := http.Get("https://www.cbr.ru/currency_base/daily/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Инициализируем goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var currencies []Currency

	currencyRUB := Currency{
		digitalCode: 643,
		litterCode:  "RUB",
		Units:       1,
		Name:        "Российский рубль",
		Rate:        1,
	}

	currencies = append(currencies, currencyRUB)

	doc.Find("tbody").Each(func(j int, s *goquery.Selection) {
		infoCurrency := s.Find("tr")
		oldlines := strings.Split(infoCurrency.Text(), "\n")
		var lines []string

		for _, value := range oldlines {
			trimmedValue := strings.TrimSpace(value)
			if trimmedValue != "" {
				lines = append(lines, trimmedValue)
			}
		}

		for i := 5; i < len(lines); i += 5 {
			digitalCode, _ := strconv.Atoi(strings.TrimSpace(lines[i]))
			litterCode := strings.TrimSpace(lines[i+1])
			units, _ := strconv.Atoi(strings.TrimSpace(lines[i+2]))
			name := strings.TrimSpace(lines[i+3])
			rateStr := strings.ReplaceAll(strings.TrimSpace(lines[i+4]), ",", ".")
			rate, _ := strconv.ParseFloat(rateStr, 32)

			currency := Currency{
				digitalCode: digitalCode,
				litterCode:  litterCode,
				Units:       units,
				Name:        name,
				Rate:        float32(rate),
			}

			currencies = append(currencies, currency)
		}
	})

	return currencies
}
