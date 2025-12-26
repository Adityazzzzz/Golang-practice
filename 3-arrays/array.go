package main

import "fmt"

func main() {
	var nums[4]int //all default as 0
	nums[0]=1
	fmt.Println(len(nums))

	nums1 := [3]int{1,2,3} //1D 
	fmt.Println(nums1)

	nums2 := [2][2]int{{1,1},{2,3}} //2D
	fmt.Println(nums2)
}