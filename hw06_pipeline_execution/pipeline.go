package main

import (
	"fmt"
	"strconv"
	"time"
)

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

const (
	sleepPerStage1 = time.Millisecond * 100
	fault1         = sleepPerStage1 / 2
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	resChan := make(Bi, 5)
	defer close(resChan)

	return resChan
	//resChan <- <-stages[3]((stages[2](stages[1](stages[0](in)))))
}

func main() {
	// Stage generator
	g := func(_ string, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			go func() {
				defer close(out)
				for v := range in {
					time.Sleep(sleepPerStage1)
					out <- f(v)
				}
			}()
			return out
		}
	}

	stages := []Stage{
		//g("Dummy", func(v interface{}) interface{} { return v }),
		//g("Multiplier (* 2)", func(v interface{}) interface{} { return v.(int) * 2 }),
		//g("Adder (+ 100)", func(v interface{}) interface{} { return v.(int) + 100 }),
		g("Stringifier", func(v interface{}) interface{} { return strconv.Itoa(v.(int)) }),
	}

	in := make(Bi, 5)
	data := []int{1, 2, 3, 4, 5}

	go func() {
		for _, v := range data {
			in <- v
		}
		close(in)
	}()

	result := make([]string, 0, 10)
	for s := range ExecutePipeline(in, nil, stages...) {
		result = append(result, s.(string))
	}

	fmt.Println(result)

}
