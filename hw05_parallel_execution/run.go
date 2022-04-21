package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Task func() error

var ErrErrorsLimitExceeded = errors.New("error limit exceeded")

func Run(tasks []Task, n, m int) error {
	// taskChan - канал с задачами Task.
	// result - результирующая ошибка функции Run.
	// gotErr, maxErr - кол-во полученных ошибок и установленный лимит ошибок.
	taskChan := make(chan Task, len(tasks))
	var result error
	var gotErr int32
	maxErr := int32(m)
	// Если m<=0, устанавливаем недостижимый лимит ошибок.
	if m <= 0 {
		maxErr = int32(len(tasks) + 1)
	}

	wg := &sync.WaitGroup{}
	// Producer. Горутина отправляет задачи из массива tasks в канал TaskChan.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, t := range tasks {
			taskChan <- t
		}
		close(taskChan)
	}()

	// Consumers. Горутины читают и выолняют задачи из канала taskChan
	// Возвращают ErrErrorsLimitExceeded при превышении лимитиа ошибок.
	wg.Add(n)
	for j := 1; j <= n; j++ {
		go func() {
			defer wg.Done()
			for gotErr < maxErr {
				t, ok := <-taskChan
				if !ok {
					return
				}
				err := t()
				if err != nil {
					atomic.AddInt32(&gotErr, 1)
				}
			}
			result = ErrErrorsLimitExceeded
		}()
	}
	// Ждем завершения горутин и возвращаем результат.
	wg.Wait()
	return result
}
