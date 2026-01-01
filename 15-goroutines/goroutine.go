package main

/*
Goroutines are very lightweight functions that run concurrently (at the same time) with other code in a Go program. They are cheaper than OS threads, so you can easily run thousands or more at once.

One-line idea:
	Add the go keyword before a function call (like go doWork()) to run that function in the background while the rest of your program keeps going.
Tiny mental model:
	Think of a goroutine as a super-light background task managed by Go’s runtime, not by the operating system, which makes concurrency simple and efficient.
*/

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}