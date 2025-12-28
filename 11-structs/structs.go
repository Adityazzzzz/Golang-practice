package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main(){
	myorder := order{
		id:"1",
		amount: 50.00,
		status: "received",
	}
	myorder.createdAt = time.Now()
	
	fmt.Println(myorder)
	fmt.Println(myorder.amount)
	fmt.Println(myorder.id)
}