package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i := 0; i < 100; i++{
		fmt.Println(from, ":", i)
	}
}

func main(){
	f("direct")

	//to call a goroutine, a light weight thread, use the following
	go f("goroutine")

	
	go func(msg string){
		fmt.Println(msg)
	}("going")



	
	go f("***************")

	
	time.Sleep(time.Second)
	fmt.Println("done")
}