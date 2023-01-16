/*
channels take in data and give out data simultaneously
if you send data into a channel, then recv it simultaneously using another thread
*/

package main

import "fmt"

func main(){

	//create a channel with chan keyword
	messages := make(chan string)

	//invoke an anonymous function in a goroutine

	//now send the message to the goroutine with the help of channel
	go func(){
		messages <- "ping"
	}()

	msg := <-messages

	fmt.Println(msg)
}