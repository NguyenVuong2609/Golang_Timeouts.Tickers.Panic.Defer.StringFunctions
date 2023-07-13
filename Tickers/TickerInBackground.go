package Tickers

import (
	"fmt"
	"time"
)

func inBackground() {
	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		fmt.Println("Ticking..")
	}
}

func TickerInBackground() {
	fmt.Println("Starting the ticker")
	go inBackground()
	fmt.Println("After goroutine..")
	time.Sleep(5 * time.Second)
	//select {}
}
