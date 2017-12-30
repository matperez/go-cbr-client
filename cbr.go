package cbr

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
)

// Currency is a currency item
type Currency struct {
	ID       string `xml:"ID,attr"`
	NumCode  uint   `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nom      uint   `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

// Result is a result representation
type Result struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Currencies []Currency `xml:"Valute"`
}

const baseURL = "http://www.cbr.ru/scripts/XML_daily.asp"
const dateFormat = "02/01/2006"

//GetCurrencyRate returns currency rate at a specific time
func GetCurrencyRate(currency string, t time.Time) (float64, error) {
	log.Printf("Fetching the currency rate for %s at %v\n", currency, t.Format("02.01.2006"))
	var result Result
	err := GetCurrencies(&result, t)
	if err != nil {
		return 0, err
	}
	for _, v := range result.Currencies {
		if v.CharCode == currency {
			return getCurrencyRateValue(v)
		}
	}
	return 0, fmt.Errorf("Unknown currency: %s", currency)
}

func getCurrencyRateValue(v Currency) (float64, error) {
	properFormattedValue := strings.Replace(v.Value, ",", ".", -1)
	floatValue, err := strconv.ParseFloat(properFormattedValue, 64)
	if err != nil {
		return 0, err
	}
	return floatValue / float64(v.Nom), nil
}

// GetCurrencies returns plain server response at a specific date
func GetCurrencies(v *Result, t time.Time) error {
	url := baseURL + "?date_req=" + t.Format(dateFormat)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("Unknown charset: %s", charset)
		}
	}
	err = decoder.Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
