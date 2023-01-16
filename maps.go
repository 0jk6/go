package main

import "fmt"

func main(){
	
	//to build a map, use the make() method
	m := make(map[string]int) //this stores {string, int} as map

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println("map:",m)

	v1 := m["key1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	//delete a key-value pair
	delete(m, "key2")

	fmt.Println(m)

	_, prs := m["key2"]
	fmt.Println("prs:", prs)

	//direct declaration
	n := map[string]int{"foo" : 1, "bar" : 2}
	fmt.Println("n:", n)
}