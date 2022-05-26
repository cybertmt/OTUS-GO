package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"unicode"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envList := make(map[string]EnvValue)
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var b []byte
	var fN string
	var enV EnvValue
	for _, file := range files {
		fN = file.Name()
		readFile, err := os.Open(dir + "/" + fN)
		defer func() { readFile.Close() }()
		if err != nil {
			return nil, err
		}
		rd := bufio.NewReader(readFile)
		s, _ := rd.ReadString('\n')
		if len(s) == 0 {
			enV.NeedRemove = true
		}
		s = strings.TrimRightFunc(s, func(c rune) bool {
			return unicode.IsSpace(c)
		})
		b = bytes.Replace([]byte(s), []byte{0}, []byte{10}, -1)
		enV.Value = string(b)
		envList[fN] = enV
	}
	return envList, err
}
