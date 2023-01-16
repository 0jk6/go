package main

import "fmt"

//declare a struct
type person struct{
	name string
	age int
}

//function to create a structure
//takes in name as string and returns the pointer to the created struct
func newPerson(name string) *person {

	//create a struct
	p := person{name : name}
	p.age = 42 //assigning the age

	return &p //return the address of the struct
}

func main(){
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age : 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age : 40})

	fmt.Println(newPerson("John"))

	s := person{name : "Sean", age : 35}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(sp.age)

	fmt.Println(s)
}