package Tickers

import (
	"fmt"
	"time"
)

func Tickers() {
	fmt.Println("Starting the ticker")

	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("Ticking..")
	}
}
