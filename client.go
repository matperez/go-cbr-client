package cbr

import (
	"net/http"
	"time"
)

// fetchFunction is a function that mimics http.Get() method
type fetchFunction func(url string) (resp *http.Response, err error)

// Client is a currency rates service client... what else?
type Client interface {
	GetRate(string, time.Time) (float64, error)
	SetFetchFunction(fetchFunction)
}

type client struct {
	fetch fetchFunction
}

func (s client) GetRate(currency string, t time.Time) (float64, error) {
	rate, err := getRate(currency, t, s.fetch)
	if err != nil {
		return 0, err
	}
	return rate, nil
}

func (s client) SetFetchFunction(f fetchFunction) {
	s.fetch = f
}

// NewClient creates a new rates service instance
func NewClient() Client {
	return client{http.Get}
}
