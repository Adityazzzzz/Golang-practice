package main

/*
A mutex (from sync.Mutex) is a simple lock that protects shared data from being accessed by multiple goroutines at once, preventing data races.

One-line idea
	Use mu.Lock() before touching shared vars, mu.Unlock() after—only one goroutine enters at a time; others wait.
Tiny example
	Protect a counter: mu.Lock(); count++; mu.Unlock() (or defer mu.Unlock()); multiple goroutines increment safely without races.
*/

// TO AVOID RACE-CONDITIONS

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock() // ***Imp
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}
	wg.Wait()

	fmt.Println(myPost.views)
}