package main

import "fmt"

func main() {
//1
	//creating map
	mpp := make(map[int]string)
	mpp[0] = "golang"
	mpp[1] = "go"
	
	fmt.Println(mpp[0])

	//delete element
	delete(mpp,1) // it will remove mpp[1]="go"


//2
	mpp2 := map[string]string{}
	mpp3 := map[string]string{"abc":"pqr", "xyz":"pqr2"}
	fmt.Println(mpp2)
	fmt.Println(mpp3)


//3 
	//check if exists or not
	_,ok := mpp3["abc"]
	if ok==true {
		fmt.Println("Exists")
	}else {
		fmt.Println("Not exists")
	}
}