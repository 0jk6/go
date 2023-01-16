package main

import (
	"fmt"
	"time"
	"sync"
)


//following function waits form one second before printing the number
func doWork() int{
	time.Sleep(time.Second)
	return 12
}


func main(){

	dataChannel := make(chan int)

	go func(){

		wg := sync.WaitGroup{}

		for i:=0; i<1000; i++ {
			
			wg.Add(1)

			go func(){
				defer wg.Done()
				result :=doWork()
				dataChannel <- result
			}()
		}

		wg.Wait()
		close(dataChannel)
	}()

	for n := range dataChannel {
		fmt.Printf("n = %d\n", n)
	}
}