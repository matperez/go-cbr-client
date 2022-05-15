package cbr

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	client := NewClient()
	rate, _ := client.GetRate("USD", time.Now())
	fmt.Printf("rate: %.2f\n", rate)
}
