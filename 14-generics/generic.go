package main

import "fmt"

/*
	Generics in Go let you write one function or type that works with many different data types, while still being fully type-safe and checked at compile time.​

One-line idea:
	Generics = functions and structs with a type parameter (like [T any]) so the same code works for int, string, etc. without duplication.​

Tiny example:
	func Print[T any](v T){
		fmt.Println(v)
	}
	can print ints, strings, or any other type using one implementation.​

	You can call it as Print(10) or Print("hello") and Go figures out the type automatically.​
*/

// polymorphism = generics


// func printSlice(arr[]int){
// 	for _,it:=range arr{
// 		fmt.Println(it)
// 	}
// }

func printSlice[T any](arr[]T){
	for _,it:=range arr{
		fmt.Println(it)
	}
}

func main(){
	printSlice([]int{1,2,3,4});
}