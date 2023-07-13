package Defer

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}
func ExampleArguments() {
	a := 5
	defer printA(a)
	a = 10
	fmt.Println("value of a before deferred function call", a)

}
