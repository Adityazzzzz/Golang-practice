package main

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