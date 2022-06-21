package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	release   = "0.1"
	buildDate = "2022-06-21"
	gitHash   = "a0b0ce55c39e112faacda929b9682f28516fdb1d"
)

func printVersion() {
	if err := json.NewEncoder(os.Stdout).Encode(struct {
		Release   string
		BuildDate string
		GitHash   string
	}{
		Release:   release,
		BuildDate: buildDate,
		GitHash:   gitHash,
	}); err != nil {
		fmt.Printf("error while decode version info: %v\n", err)
	}
}
