package main

import "fmt"


func zeroval(int_val int){
	int_val = 0
}

func zeroptr(int_ptr *int){
	*int_ptr = 0
}


func main(){
	i := 1

	fmt.Println("initial value:", i)
	
	zeroval(i)
	fmt.Println("after zeroval():", i)

	zeroptr(&i)
	fmt.Println("after zeroptr():", i)

	fmt.Println("pointer:", &i)
}