package main

import "fmt"


func main(){
	//assign a slice with make fun

	s := make([]string, 3)
	fmt.Println("emp:", s)

	//assing values
	s[0] = "hi"
	s[1] = "wait"
	s[2] = "what"

	fmt.Println("set:", s)

	
	//s[3] = "what"; this does not work, use append method to add new element
	s = append(s, "the")
	s = append(s, "fuck", "?")

	fmt.Println(s)

	fmt.Println("get:", s[4])


	//we can copy the slice btw
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)


	l := s[2:5]
	fmt.Println("l:", l)
}