package main

import "fmt"

func main(){
	messages := make(chan string, 2) //make a channel of strings buffering up to 2 values

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}