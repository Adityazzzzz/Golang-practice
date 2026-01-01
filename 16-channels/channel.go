package main

import (
	"fmt"
	"time"
)

/*
Channels in Go are typed pipes that let goroutines safely send and receive data, making communication between concurrent code simple and race-free.

One-line idea:
	Create with ch := make(chan int); send ch <- 5, receive x := <-ch—blocks until both sides are ready, syncing goroutines too.

Tiny example:
	One goroutine does ch <- "ping"; main does msg := <-ch to get it—data flows safely without shared memory issues.
*/



func sample(){
	messageChannel := make(chan string)

//send:
	messageChannel <- "ping"

//receive:
	msg := <- messageChannel
	fmt.Println(msg)
}



func processChan(numChannel chan int){
	fmt.Println("processing num", <- numChannel)
}
func main(){
	numChannel := make(chan int)
	go processChan(numChannel)

	numChannel<-5

	time.Sleep(time.Second*2)
}

