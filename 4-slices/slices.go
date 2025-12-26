package main

import (
	"fmt"
	"slices"
)

/*
	Slice -> dynamic
	mostly used in go, useful methods.
*/

func main(){
	//(a) uninitialised
	var nums[]int
	//(b)
	var nums1 = make([]int,2,5) //size=2, capacity=5..[0,0]
	nums1 = append(nums1, 1) // [0,0,1]
	nums1 = append(nums1, 2) // [0,0,1,1]
	//(c)
	arr := []int{}

	//2.For copy
	const n = 5
	var arr1=make([]int,n)
	var arr2 = make([]int, len(arr1))
	copy(arr2,arr1)

	//3.slice operator
	var nums3 = []int{1,2,3}
	fmt.Println(nums3[0:2])

	//4.compare
	fmt.Println(slices.Equal(arr1,arr2))//return true/false

	//2D
	twoDarr := [][]int{{0,0},{1,2}}

}