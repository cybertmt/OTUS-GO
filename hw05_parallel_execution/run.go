package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(id int, tasksChan <-chan Task, errChan chan<- error, resChan chan<- bool) {
	// for t := range tasksChan {
	// 	x := t
	// 	fmt.Printf("Worker id %d got %v\n", id, x())
	// 	time.Sleep(time.Millisecond)
	// 	errChan <- t()
	// }
	for {
		v, ok := <-tasksChan
		if !ok {
			resChan <- true
		}
		x := v
		fmt.Printf("Worker id %d got %v\n", id, x())
		time.Sleep(time.Second)
		errChan <- v()
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksChan := make(chan Task, len(tasks))
	errorChan := make(chan error, m)
	resultChan := make(chan bool, 1)
	for i := 1; i <= n; i++ {
		i := i
		go worker(i, tasksChan, errorChan, resultChan)
	}

	for _, t := range tasks {
		// x := t
		// fmt.Printf("Loading %v\n", x())
		tasksChan <- t
	}
	close(tasksChan)

	go func() {
		for j := 0; j < m; j++ {
			e := <-errorChan
			fmt.Printf("Got %v\n", e)

		}
		// 	resultChan <- false
	}()
	// <-resultChan
	// return nil
	res := <-resultChan
	if res {
		return nil
	}
	return ErrErrorsLimitExceeded
}

func main() {
	tasksCount := 50
	tasks := make([]Task, 0, tasksCount)

	var runTasksCount int32

	for i := 0; i < tasksCount; i++ {
		err := fmt.Errorf("error from task %d", i)
		tasks = append(tasks, func() error {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
			atomic.AddInt32(&runTasksCount, 1)
			return err
		})
	}
	fmt.Println(Run(tasks, 5, 20))
}
