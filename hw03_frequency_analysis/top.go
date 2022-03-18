package main

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	// Для пустого запроса возвращаем nil слайс
	if str == "" {
		return []string{}
	}
	// Слайс слов, возвращаемый функцией Top10
	strResult := make([]string, 0)
	// Map для подсчета слов в тексте [слово]количество
	wordMap := make(map[string]int)
	// Структура для подсчета количества слов,
	// используется для сортировки по значениям
	type wordStruct struct {
		Word  string
		Count int
	}
	// Разбираем исходную строку на слайс слов с пробелом в качестве разделителя
	strSlice := strings.Fields(str)
	// Наполняем map с подсчетом количества слов
	for _, s := range strSlice {
		value, ok := wordMap[s]
		if ok {
			wordMap[s] = value + 1
			continue
		}
		wordMap[s] = 1
	}
	// Создаем слайс структур для слов и наполняем его из map с количеством слов
	wordStructSlice := make([]wordStruct, 0)
	for k, v := range wordMap {
		wordStructSlice = append(wordStructSlice, wordStruct{k, v})
	}
	// Сортируем полученный слайс по количеству, затем равные значения лексикографически
	sort.Slice(wordStructSlice, func(i, j int) bool {
		if wordStructSlice[i].Count != wordStructSlice[j].Count {
			return wordStructSlice[i].Count > wordStructSlice[j].Count
		}
		return wordStructSlice[i].Word < wordStructSlice[j].Word
	})
	// Первые 10 слов отсортированного слайса добавляем в результирующий слайс и возвращаем
	for _, v := range wordStructSlice[:10] {
		strResult = append(strResult, v.Word)
	}
	return strResult
}
