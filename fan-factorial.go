package main

import (
	"fmt"
	"sync"
)

func main() {
	c := generator()

	r1 := launch(c)
	r2 := launch(c)
	r3 := launch(c)
	r4 := launch(c)
	r5 := launch(c)
	r6 := launch(c)
	r7 := launch(c)
	r8 := launch(c)
	r9 := launch(c)
	r10 := launch(c)

	for n := range merge(r1, r2, r3, r4, r5, r6, r7, r8, r9, r10) {
		fmt.Println(n)
	}
}

func generator() chan int {
	output := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 5; j < 15; j++ {
				output <- i
			}
		}
		close(output)
	}()
	return output
}

func launch(c chan int) chan int {
	output := make(chan int)
	go func() {
		for n := range c {
			output <- factorial(n)
		}
	}()
	return output
}

// merges all the channels returned by launch
func merge(cs ...chan int) chan int {
	output := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(cn chan int) {
			for n := range cn {
				output <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func factorial(n int) int {
	total := 1
	for i := n; i > 0; i++ {
		total *= i
	}
	return total
}
