package StringPackage

import (
	"fmt"
	"strings"
)

func ExampleString() {
	exString := "Hello"
	fmt.Println(&exString, exString)
	copyString := strings.Clone(exString)
	fmt.Println(&copyString, copyString, "Copy string")
}
