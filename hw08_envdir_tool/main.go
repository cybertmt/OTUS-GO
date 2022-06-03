package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	// Принимаем аргуменны командной строки.
	args := os.Args
	// Проверяем необходимое кол-во аргументов.
	if len(args) < 3 {
		log.Fatal("Not enough arguments. Usage: go-envdir /path/to/evndir command arg1 arg2...")
	}
	// Читаем переменные окружения из директории, указанной в аргементе 1.
	env, err := ReadDir(args[1])
	if err != nil {
		if !errors.Is(err, ErrWrongFileName) {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
	// Запускаем внешнюю программу с переменныи командной строки и map переменных окружения.
	exitCode := RunCmd(args[2:], env)
	os.Exit(exitCode)
}
