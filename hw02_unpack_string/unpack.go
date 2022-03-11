package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Unpack Распаковка строки по ТЗ
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
	result, sVal := "", ""
	var iVal int
	var err error

	// Перебор символов строки
	for i, val := range runes {

		// Если руна является цифрой
		if unicode.IsDigit(val) {

			// Переведение необходимых типов
			sVal = string(val)
			iVal, err = strconv.Atoi(sVal)

			if err != nil {
				return "", ErrInvalidString
			}

			// Если перед цифрой один или несколько символов `\`
			if i > 0 && string(runes[i-1]) == `\` {
				if i > 1 && string(runes[i-2]) == `\` {
					if i > 2 && string(runes[i-3]) == `\` {
						result = result[:len(result)-3]
						result += `\` + sVal
						continue
					}
					result += strings.Repeat(`\`, iVal-2)
					continue
				}
				result = result[:len(result)-1]
				result += string(val)
				continue
			}

			// Если после цифры идет еще одна цифра
			if len(runes)-1 > i && unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}

			// Если значение цифры 0
			if iVal == 0 {
				result = result[:len(result)-1]
				continue
			}

			// Повторить символ перед цифрой указанное число раз
			result += strings.Repeat(string(runes[i-1]), iVal-1)
			continue
		}

		// Добавить символ в результирующую строку
		result += string(val)

	}
	return result, nil
}

func main() {
	s := `&^%`
	result, err := Unpack(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
