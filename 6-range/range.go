package main

import "fmt"

/*
	Used for iterating over data structures
*/
func main(){
	arr := []int{1,2,3}
//simple for loop
	for i:=0;i<len(arr);i++ {
		fmt.Println(arr[i])
	}

// using range
	sum :=0
	for idx,it := range arr {
		sum = sum+it
		fmt.Println(idx,it)
	}
}