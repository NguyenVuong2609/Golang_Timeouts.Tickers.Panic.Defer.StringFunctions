package Defer

import "fmt"

type dog struct {
	name string
}
type Sound interface {
	makeSound()
}

func (d dog) makeSound() {
	fmt.Println("Name:", d.name, "Gou gou!")
}

func ExampleDeferMethod() {
	d := dog{"Milu"}
	defer d.makeSound()
	fmt.Printf("Welcome ")
}
