package main

import "fmt"

//anything that takes a receiver is a method in go

type rect struct {
	width, height int
}

//pointer is passed without a copy
func (r *rect) area() int {
	return r.width * r.height
}

//a copy of the struct is passed here
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main(){
	r := rect{width:10, height:5}

	//above defined methods can be accessed from the above defined variable r as r.area() and r.perimeter()
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())

}