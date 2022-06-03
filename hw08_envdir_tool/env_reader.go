package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ErrWrongFileName ошибка в имени файла, содержащем '='.
var ErrWrongFileName = fmt.Errorf("some env files have wrong names: include '=', skipped")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	// Создаем map переменных окружения.
	envList := make(Environment)
	// Читаем имена файлов в директории dir в массив files.
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	// fN - filename, enV - объект значения переменной окружения.
	var fN string
	var enV EnvValue
	// Проходим по списку имен файлов.
	for _, file := range files {
		// Не является ли файл директорией?
		if file.IsDir() {
			continue
		}
		fN = file.Name()
		// Если имя файла содержит '=', пропускаем файл и информируем в ошибке.
		if strings.Contains(fN, "=") {
			err = ErrWrongFileName
			continue
		}
		// Открываем файл.
		readFile, err := os.Open(filepath.Join(dir, fN))
		defer func() { readFile.Close() }()
		if err != nil {
			return nil, err
		}
		// Читаем из файла первую строку.
		rd := bufio.NewReader(readFile)
		s, err := rd.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
		// Если файл пустой, маркируем переменную под удаление.
		if len(s) == 0 {
			enV.NeedRemove = true
		}
		// Удаляем из строки табуляцию и пробелы справа.
		s = strings.TrimRight(s, " \t\n")
		// Null-байты заменяем на '\n'.
		enV.Value = string(bytes.ReplaceAll([]byte(s), []byte{0}, []byte{10}))
		// Добавляем готовую переменную окружения в map.
		envList[fN] = enV
	}
	// Возвращаем map с переменными окружения и ошибку.
	return envList, err
}
