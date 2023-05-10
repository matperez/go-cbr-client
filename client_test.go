package cbr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetRate(t *testing.T) {
	client := NewClient()
	rate, err := client.GetRate("USD", time.Now())
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, rate, float64(1))
}

func TestGetRate_Error(t *testing.T) {
	client := NewClient()
	rate, err := client.GetRate(" ", time.Now())
	assert.Error(t, err)
	assert.Equal(t, float64(0), rate)
}
