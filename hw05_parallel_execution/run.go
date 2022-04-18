package main

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// worker обработчик заданий Task.
func worker(wg *sync.WaitGroup, taskChan <-chan Task, doneChan <-chan bool, errChan chan<- error) {
	defer wg.Done()
	for {
		// Останавливаемся по сигналу <-doneChan
		// или когда закончатся задания.
		select {
		case <-doneChan:
			return
		default:
			t, ok := <-taskChan
			if !ok {
				return
			}
			e := t()
			if e != nil {
				errChan <- e
			}
		}
	}
}

// Run запускает задачи в n горутинах и останавливает их работу при получении m ошибок.
func Run(tasks []Task, n, m int) error {
	// taskChan - канал задач.
	// errChan - канал ошибок выполнения задач.
	// resChan - канал с возвращаемым значением.
	// doneChan - канал для остановки работы горутин.
	taskChan := make(chan Task, n+m)
	errChan := make(chan error, n+m)
	resChan := make(chan error, 1)
	doneChan := make(chan bool)
	wg := &sync.WaitGroup{}

	// Consumer worker pool.
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(wg, taskChan, doneChan, errChan)
	}

	// Producer - отправляет задачи в канал taskChan.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, t := range tasks {
			taskChan <- t
		}
		close(taskChan)
	}()

	// Горутина - счетчик ошибок.
	wg2 := &sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		defer close(resChan)
		defer close(doneChan)
		// Если m <=0, ошибки не считаем.
		if m <= 0 {
			for {
				_, ok := <-errChan
				if !ok {
					resChan <- nil
					return
				}
			}
		}
		// Считаем ошибки.
		// Если ошибки не превысили порог, отправляем nil в канал результата.
		for j := 0; j < m; j++ {
			_, ok := <-errChan
			if !ok {
				resChan <- nil
				return
			}
		}
		// При превышении порога ошибок (m) останавливаем горутины
		// и отправляем ошибку в канал результата.
		resChan <- ErrErrorsLimitExceeded
	}()
	// Ждем завершения worker pool и горутины заданий.
	wg.Wait()
	// Закрываем канал ошибок и ждем завершения горутины подсчета ошибок.
	close(errChan)
	wg2.Wait()

	return <-resChan
}
