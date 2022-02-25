package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Системный вывод перевернутой строки "Hello, OTUS!"
	fmt.Println(stringutil.Reverse("Hello, OTUS!"))
}
