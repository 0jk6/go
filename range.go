package main

import "fmt"

func main(){
	
	nums := []int{2,3,4}
	sum := 0



	for _, num := range nums{
		sum += num
	}

	fmt.Println("sum:", sum)

	mp := map[string]string{"a" : "apple", "b" : "banana"}

	for key, value := range mp{
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range mp{
		fmt.Println("key:", key)
	}

	//similarly we can iterate over a string too
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}