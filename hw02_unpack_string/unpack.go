package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	// Проверка пустой строки
	if s == "" {
		return "", nil
	}
	// Конвертация строки в слайс рун
	runes := []rune(s)
	// Проверка первого символа строки: является ли цифрой
	if unicode.IsDigit(runes[0]) {
		return "", ErrInvalidString
	}

	// result - возвращаемая и удачно распакованная строка
	// sVal - символ строки, сконвертированный в тип string
	// iVal - символ строки, сконвертированный в тип int
	// err - ошибка при конвертациях и приведениях к типу
	var result, sVal string
	var iVal int
	var err error

	// Перебор символов строки
	for i := 0; i < len(runes); {
		// i-ая руна
		val := runes[i]
		sVal = string(val)
		// Если руна является симовлом `\` и после нее есть еще руны
		if sVal == `\` && len(runes)-1 > i {
			// Добавляем в результирующую строку следующий за `\` экранированный символ
			result += string(runes[i+1])
			i += 2
			continue
		}
		// Если руна является цифрой
		if unicode.IsDigit(val) {
			// Преведение необходимых типов
			iVal, err = strconv.Atoi(sVal)

			if err != nil {
				return "", ErrInvalidString
			}
			// Если после цифры идет еще одна цифра, вернуть ошибку
			if len(runes)-1 > i && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
			// Повторить символ перед цифрой указанное число раз
			result = result[:len(result)-1] + strings.Repeat(string(runes[i-1]), iVal)
			i++
			continue
		}
		// Добавить символ в результирующую строку
		result += string(val)
		i++
	}
	return result, nil
}

func main() {
	s := `\ad5e\\5\7`
	result, err := Unpack(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
