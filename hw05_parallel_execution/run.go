package main

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskChan := make(chan Task, len(tasks))
	errChan := make(chan error, n)
	resChan := make(chan error)
	doneChan := make(chan bool)
	wg := &sync.WaitGroup{}

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func() {
			// fmt.Printf("Worker %v started\n", i)
			defer wg.Done()
			for {
				select {
				case <-doneChan:
					// fmt.Printf("Worker %v finished cause of error, stopping\n", i)
					return
				default:
					t, ok := <-taskChan
					if !ok {
						// fmt.Printf("Worker %v got nothing to do, stopping\n", i)
						// close(errChan)
						// resChan <- nil
						return
					}
					e := t()
					// fmt.Printf("Worker %v got %v\n", i, e)
					if e != nil {
						errChan <- e
						continue
					}
					// fmt.Printf("Worker %v got %v\n", i, e)
					continue
				}
			}
		}()
	}

	go func() {
		for _, t := range tasks {
			taskChan <- t
		}
		close(taskChan)
	}()

	go func() {
		for j := 0; j < m; j++ {
			_, ok := <-errChan
			if !ok {
				// fmt.Printf("End of tasks because of %v\n", ok)
				close(doneChan)
				close(resChan)
				return
			}
			// fmt.Println("Catch error!", e)
		}
		// fmt.Println("Too much errors!")
		// for z := 0; z < n; z++ {
		//	doneChan <- true
		// }

		close(doneChan)
		// close(errChan)
		resChan <- ErrErrorsLimitExceeded
		close(resChan)
	}()
	// close(resChan)

	wg.Wait()
	close(errChan)
	// fmt.Println("Closing ErrChan")

	// for res := range resChan {
	//	if res != nil {
	//		return res
	//	}
	// }
	// for res := range resChan {
	//	if res != nil {
	//		return res
	//	}
	// }

	// result := <-resChan
	// fmt.Println("Got result:", result)

	return <-resChan
}

//  func main() {
//	b := true
//	//b = false
//	var tasksCount, n, m int
//	tasks := make([]Task, 0, tasksCount)
//	if b {
//		tasksCount = 20
//		n = 50
//		m = 19
//		tasks = make([]Task, 0, tasksCount)
//
//		var runTasksCount int32
//
//		for i := 0; i < tasksCount; i++ {
//			err := fmt.Errorf("error from task %d", i)
//			tasks = append(tasks, func() error {
//				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
//				atomic.AddInt32(&runTasksCount, 1)
//				return err
//			})
//		}
//
//	} else {
//		tasksCount = 20
//		n = 10
//		m = 1
//		tasks = make([]Task, 0, tasksCount)
//
//		var runTasksCount int32
//		var sumTime time.Duration
//
//		for i := 0; i < tasksCount; i++ {
//			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
//			sumTime += taskSleep
//
//			tasks = append(tasks, func() error {
//				time.Sleep(taskSleep)
//				atomic.AddInt32(&runTasksCount, 1)
//				return nil
//			})
//		}
//	}
//
//	res := Run(tasks, n, m)
//	if res != nil {
//		fmt.Println(res)
//		return
//	}
//	fmt.Println("Work Complete")
// }
