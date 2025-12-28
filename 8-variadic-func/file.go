package main

import "fmt"

//how to make 
func sum(arr ...int)int{
	total:=0
	for _,it:=range arr{
		total = total+it;
	}
	return  total
}

func main() {
	ans := sum(1,2,3,4,5,56)
	fmt.Println(ans)
	//or
	arr:=[]int{1,2,3,4,4,5,67,8,9}
	res:=sum(arr...)
	fmt.Println(res)
}