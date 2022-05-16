package main

import (
	"fmt"
	"time"

	cbr "github.com/ivanglie/go-cbr-client"
)

func main() {
	client := cbr.NewClient()
	rate, err := client.GetRate("USD", time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(rate)
}
