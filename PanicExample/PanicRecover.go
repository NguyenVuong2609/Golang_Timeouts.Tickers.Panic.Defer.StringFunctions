package PanicExample

import "fmt"

var n = []int{1, 2, 3}

func slicePanic() {
	//defer recoverFromPanic()
	fmt.Println(n[5])
	fmt.Println("returned from slice")
}
func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recover successful!", r)
	}
	fmt.Println("I'm Professor!")
}
func nextDay() {
	fmt.Println("Thank you!")
	fmt.Println(n[3])
}
func PanicRecover() {
	defer recoverFromPanic()
	defer nextDay()
	slicePanic()
	fmt.Println("return main")
}
func FatherPanic() {
	PanicRecover()
	fmt.Println("Last chance!")
}
