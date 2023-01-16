package main

import "fmt"

func main(){
	var a [5]int

	fmt.Println(a)

	b := [5]int{1,2,3,4,5}

	fmt.Println(b)

	var c [5]int

	for i:=0; i<5; i++{
		c[i] = i*10
	}

	fmt.Println(c)


	//2D arrays
	var td [2][3]string

	for i:=0; i<2; i++{
		for j:=0; j<3; j++{
			td[i][j] = "*"
		}
	}

	fmt.Println(td)
}