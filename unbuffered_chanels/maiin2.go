package main

import (
	"fmt"
	"time"
)

func main2() {
	myChanel := make(chan int)

	go func() {

		for i := 1; i <= 10; i++ {
			myChanel <- i
			fmt.Println("Sent Data : ", i)
			time.Sleep(1 * time.Second)
		}
		close(myChanel)
	}()

	for {
		data, isOpen := <-myChanel
		if !isOpen {
			break
		}
		fmt.Println("Recived Data : ", data)

	}

}
