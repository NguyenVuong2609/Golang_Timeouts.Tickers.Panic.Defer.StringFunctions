package Timeouts

import (
	"fmt"
	"time"
)

func timeConsuming() string {
	time.Sleep(5 * time.Second)
	return "The timeConsuming() function has stopped"
}

func Timeouts() {
	currentChannel := make(chan string, 1)

	go func() {
		text := timeConsuming()
		currentChannel <- text
	}()

	select {
	case res := <-currentChannel:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Out of time :(")
	}
	fmt.Println("Main function exited!")
}
