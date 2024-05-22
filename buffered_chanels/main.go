package main

import (
	"fmt"
	"time"
)

func main1() {
	//myChanel := make(chan int, 2)

	// data := <-myChanel
	// fmt.Println(data)

	// myChanel <- 1
	// myChanel <- 2

	// fmt.Println(<-myChanel)

	// myChanel <- 3

	// fmt.Println(<-myChanel)

	message := make(chan string, 2)
	go func() {
		data1 := <-message
		fmt.Println("First :", data1)
		data2 := <-message
		fmt.Println("Second : ", data2)
	}()

	message <- "Hellooo"
	message <- "world "
	message <- "hii"

	time.Sleep(1 * time.Second)

	fmt.Println("End of the main")

}
