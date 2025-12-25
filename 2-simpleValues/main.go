package main

import "fmt"

func main() {
	//DATA TYPES
	fmt.Println(1)
	fmt.Println("string")
	fmt.Println(1+2)
	fmt.Println(true)
	fmt.Println(1.6)

	//VARIABLES : "var"
	{
		//1.
		var name string = "golang"
		//2.
		var name2 = "go"
		//3.
		name3 := "golang"

		fmt.Println(name)
		fmt.Println(name2)
		fmt.Println(name3)
	}


	//CONSTANTS : "const"
	{
		//1
		const age = 10
		fmt.Println(age)

		//2
		const (
			port =5000
		)
		fmt.Println(port)
	}

}