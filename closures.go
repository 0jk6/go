package main

import "fmt"


//used for inline functions

func integerSequence() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}


func main(){

	nextInt := integerSequence()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := integerSequence()
	fmt.Println(newInts())
	
}