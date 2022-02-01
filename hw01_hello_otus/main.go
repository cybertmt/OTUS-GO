package main

import (
	"os"

	"golang.org/x/example/stringutil"
)

func main() {
	// Выводим в стандартный вывод перевернутую строку "Hello, OTUS!"
	os.Stdout.Write([]byte(stringutil.Reverse("Hello, OTUS!")))
}
