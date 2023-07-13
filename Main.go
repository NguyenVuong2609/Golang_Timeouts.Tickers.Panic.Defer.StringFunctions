package main

import (
	"Day5-Timeouts-Tickers-Panic-Defer/StringPackage"
	"fmt"
)

func main() {
	fmt.Println("-------------------- Timeouts ------------------------")
	//Timeouts.Timeouts()
	fmt.Println("-------------------- Context Deadline ------------------------")
	//Timeouts.ContextDeadline()
	fmt.Println("-------------------- Tickers ------------------------")
	//Tickers.Tickers()
	//Tickers.Ticker2()
	//Tickers.TickerInBackground()
	fmt.Println("-------------------- Panic ------------------------")
	//PanicExample.Panic()
	//PanicExample.FatherPanic()
	//PanicExample.PanicOtherGoroutines()
	fmt.Println("-------------------- Defer ------------------------")
	//Defer.ExampleDefer()
	//Defer.ExampleDeferMethod()
	//Defer.ExampleArguments()
	fmt.Println("-------------------- String functions ------------------------")
	StringPackage.ExampleString()
}
