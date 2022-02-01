package main

import (
	"golang.org/x/example/stringutil"
	"os"
)

func main() {
	// Выводим перевернутую строку "Hello, OTUS!"
	os.Stdout.Write([]byte(stringutil.Reverse("Hello, OTUS!")))
}
