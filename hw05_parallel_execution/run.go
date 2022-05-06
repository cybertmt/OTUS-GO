package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

type Task func() error

var ErrErrorsLimitExceeded = errors.New("error limit exceeded")

// Run запускает задачи в n горутинах и останавливает их работу при получении m ошибок.
func Run(tasks []Task, n, m int) error {
	// taskChan - канал с задачами Task.
	// gotErr, maxErr - кол-во полученных ошибок и установленный лимит ошибок.
	taskChan := make(chan Task, len(tasks))
	var gotErr int32
	maxErr := int32(m)
	// Если m<=0, устанавливаем недостижимый лимит ошибок.
	if m <= 0 {
		maxErr = int32(len(tasks) + n)
	}
	wg := &sync.WaitGroup{}

	// Producer. Горутина отправляет задачи из массива tasks в канал TaskChan.
	for _, t := range tasks {
		taskChan <- t
	}
	close(taskChan)

	// Consumers. Горутины читают и выолняют задачи из канала taskChan
	// Возвращают ErrErrorsLimitExceeded при превышении лимитиа ошибок.
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for atomic.LoadInt32(&gotErr) < maxErr {
				t, ok := <-taskChan
				if !ok {
					return
				}
				err := t()
				if err != nil {
					atomic.AddInt32(&gotErr, 1)
				}
			}
		}()
	}

	// Ждем завершения горутин и возвращаем результат.
	wg.Wait()
	if gotErr > maxErr {
		return ErrErrorsLimitExceeded
	}
	return nil
}
