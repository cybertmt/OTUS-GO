package main

import (
	"os"

	"golang.org/x/example/stringutil"
)

func main() {
	// Системный вывод перевернутой строки "Hello, OTUS!"
	_, _ = os.Stdout.Write([]byte(stringutil.Reverse("Hello, OTUS!")))
}
