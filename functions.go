package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func plusPlus(a, b, c int) int {
	return a + b + c;
}


func multipleReturns() (int, int) {
	return 2, 3
}


//varadic function
func varadic(nums ...int) int {
	fmt.Println(nums)

	sum := 0

	for _ , num := range nums {
		sum += num
	}

	return sum;
}



func main(){
	res := plus(2, 3)
	fmt.Println("2+3:", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3:", res)

	a, b := multipleReturns()
	fmt.Println("a:", a, "b:", b)

	sum := varadic(1,2,3,4,5,6,7,8,9)

	fmt.Println("variadic sum: ", sum)

	nums := []int{1, 2, 3, 4}
    fmt.Println(varadic(nums...))
}