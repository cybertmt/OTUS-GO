package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Обновляем переменные окружения согласно map.
	for enVar, value := range env {
		// Удаляем переменные, помеченные для удаления.
		if value.NeedRemove {
			os.Unsetenv(enVar)
			continue
		}
		// Если переменная уже есть в окружении, удалаем её.
		if _, ok := os.LookupEnv(enVar); ok {
			os.Unsetenv(enVar)
		}
		// Устанавливаем ключи и значения переменных окружения.
		os.Setenv(enVar, value.Value)
	}
	// Отделяем имя программы и аргумены.
	name, args := cmd[0], cmd[1:]
	// Вызываем внешнюю программу.
	proc := exec.Command(name, args...)
	// Добавляем дополнительные переменные локального окружения.
	proc.Env = append(os.Environ(), args...)
	// Прокидываем потоки ввода/вывода/ошибок в приложение.
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	proc.Stdin = os.Stdin
	// В случае ошибки возвращаем exit code 1.
	if err := proc.Run(); err != nil {
		return 1
	}
	// В случае корректного исполнения возвращаем exit code 0.
	return
}
