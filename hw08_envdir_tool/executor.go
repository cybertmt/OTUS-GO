package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for enVar, value := range env {
		if value.NeedRemove {
			os.Unsetenv(enVar)
			fmt.Printf("%v=%v\n", enVar, os.Getenv(enVar))
			continue
		}
		_, ok := os.LookupEnv(enVar)
		if ok {
			os.Unsetenv(enVar)
		}
		os.Setenv(enVar, value.Value)
		fmt.Printf("%v=%v\n", enVar, os.Getenv(enVar))
	}
	cmdS := exec.Command("git", "commit", "-am", "fix")
	err := cmdS.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmdS.Wait() // ошибка выполнения
	if err != nil {
		log.Fatal(err)
	}

	return
}
