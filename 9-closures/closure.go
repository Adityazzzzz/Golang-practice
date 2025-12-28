package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count += 1
		return count
	}
}

func main() {
	inc := counter()
	fmt.Println(inc()) //1
	fmt.Println(inc()) //2
}
