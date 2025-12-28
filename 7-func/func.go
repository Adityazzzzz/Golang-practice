package main

import "fmt"

/*
basic:
	func add(a int, b int) int {
		return a + b
	}
	func main() {
		fmt.Println(add(1,2))
	}
*/

func helper() func(a int) int {
	return func(a int) int{
		return 4
	}
}

func main(){
	ans:=helper()
	ans(7)
}