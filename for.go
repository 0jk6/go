package main

import "fmt"

func main(){
	//short hand assignment
	i := 1

	//for loop with only one condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//general for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//for loop without condition
	k := 1
	for {
		fmt.Println("loop")
		
		if k >= 3 {
			break
		}

		k = k + 1
	}

	//for loop with continue statement
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}