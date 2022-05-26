package main

import (
	"fmt"
)

func main() {
	commands := []string{"arg1", "arg2"}
	dir := "testdata/env"
	env, err := ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		//return
	}
	i := RunCmd(commands, env)

	//fmt.Println(env)
	fmt.Println(i)
}
