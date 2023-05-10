package cbr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Debug(t *testing.T) {
	Debug = true
	rate("CNY", time.Now())
	assert.True(t, Debug)

	Debug = false
	rate("CNY", time.Now())
	assert.False(t, Debug)
}

func Test_rate_Error(t *testing.T) {
	rate, err := rate("", time.Now())
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), rate)
}

func Test_currencyRateValue_Error(t *testing.T) {
	c := Currency{}
	c.Value = "0'1"
	rate, err := currencyRateValue(c)
	assert.NotNil(t, err)
	assert.Equal(t, float64(0), rate)
}
