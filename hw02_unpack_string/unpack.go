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

	// result - возвращаемая удачно распакованная строка
	// sVal - символ сконвертированный в тип string
	// iVal - символ сконвертированный в тип int
	// err - ошибка при конвертациях и приведениях к типу
	var result, sVal string
	var iVal int
	var err error

	// Перебор символов строки
	for i, val := range runes {
		// Если руна является цифрой
		fmt.Println(i, string(val))
		if unicode.IsDigit(val) {
			// Преведение необходимых типов
			sVal = string(val)
			iVal, err = strconv.Atoi(sVal)

			if err != nil {
				return "", ErrInvalidString
			}

			if i > 0 && string(runes[i-1]) == `\` {
				result = result[:len(result)-1]
				result += string(val)
				continue
			}

			// Если после цифры идет еще одна цифра
			if len(runes)-1 > i && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}

			// Повторить символ перед цифрой указанное число раз
			result = result[:len(result)-1]
			result += strings.Repeat(string(runes[i-1]), iVal)
			fmt.Println(result)
			continue
		}
		if i > 0 && string(runes[i-1]) == `\` && string(runes[i]) == `\` {
			result = result[:len(result)-1]
			result += string(val)
			fmt.Println("Hit!")
			fmt.Println(result)
		}

		// Добавить символ в результирующую строку
		result += string(val)
		fmt.Println(result)
	}
	return result, nil
}

func main() {
	s := `qwe\\\3`
	result, err := Unpack(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
