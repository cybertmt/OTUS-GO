package main

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[;,&?!'.:]|\s-`)

func Top10(str string) []string {
	// Для пустого запроса возвращаем nil слайс
	if str == "" {
		return []string{}
	}
	// Для задания со '*' очищаем текст от знаков и конвертируем заглавные буквы в строчные
	str = re.ReplaceAllString(str, " ")
	str = strings.ToLower(str)
	// Разбираем исходную строку на слайс слов с пробелом в качестве разделителя
	strSlice := strings.Fields(str)
	// Если слайс остался пустой, возвращаем nil слайс
	if len(strSlice) == 0 {
		return []string{}
	}
	// Map для подсчета слов в тексте [слово]количество
	wordMap := make(map[string]int)
	// Наполняем map с подсчетом количества слов
	for _, s := range strSlice {
		wordMap[s]++
	}
	// Структура для подсчета количества слов,
	// используется при сортировке по значениям полей
	type wordStruct struct {
		Word  string
		Count int
	}
	// Создаем слайс структур для слов и наполняем его из map с количеством слов
	wordStructSlice := make([]wordStruct, 0, len(wordMap))
	for k, v := range wordMap {
		wordStructSlice = append(wordStructSlice, wordStruct{k, v})
	}
	// Сортируем полученный слайс по полю 'количество', затем по полю 'слово' лексикографически
	sort.Slice(wordStructSlice, func(i, j int) bool {
		if wordStructSlice[i].Count != wordStructSlice[j].Count {
			return wordStructSlice[i].Count > wordStructSlice[j].Count
		}
		return wordStructSlice[i].Word < wordStructSlice[j].Word
	})
	// Слайс слов, возвращаемый функцией Top10
	strResult := make([]string, 0, 10)
	// Первые 10 слов отсортированного слайса добавляем в результирующий слайс и возвращаем
	for i := 0; i < 10 && len(wordStructSlice)-i > 0; i++ {
		strResult = append(strResult, wordStructSlice[i].Word)
	}
	return strResult
}
