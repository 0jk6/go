package main

import "fmt"

func main(){

	var a = "initial"
	fmt.Println(a)
	
	//b and c have a type int
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//e will have zero value
	var e int
	fmt.Println(e)

	//short hand notation
	f := "apple"
	fmt.Println(f)
}