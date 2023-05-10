package cbr

import (
	"time"
)

// Client is a rates service client interface.
type Client interface {
	GetRate(string, time.Time) (float64, error)
}

type client struct {
}

// GetRate returns a currency rate for a given currency and date.
func (s *client) GetRate(currency string, t time.Time) (float64, error) {
	rate, err := rate(currency, t)
	if err != nil {
		return 0, err
	}
	return rate, nil
}

// NewClient creates a new rates service instance
func NewClient() Client {
	return &client{}
}
