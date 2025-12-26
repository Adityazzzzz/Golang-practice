package main

import "fmt"
import "time"

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

	//LOOPS
	{
		i:=1
		for i<=3{
			//print
			i = i+1
		}
		//or
		for i:=0;i<=3;i++ {
			//print
			continue
		}
	}

	//IF-ELSE
	{
		num:=1
		if num>0 {
			fmt.Println(num)
		}else {
			fmt.Println(num)
		}
	}

	//SWITCH
	{
		j:=1
		switch j{
			case 1: fmt.Println(1)
			case 2: 
			default:
		}
		
		// Multiple switch
		switch time.Now().Weekday(){
		case time.Thursday:
			fmt.Println("its thursday")
		}

		// Type Switch
		whoAmI := func(i interface{}){
			switch t := i.(type){
				case int:
					fmt.Println("i is int",t)
				case string:
					fmt.Println("i is string",t)
			}
		}
		whoAmI(50)
	}
}