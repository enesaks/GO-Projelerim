package main

import "fmt"

func main() {
	myChanel := make(chan int)

	go func() {
		data := <-myChanel
		fmt.Println("first routune data : ", data)
	}()
	go func() {
		data := <-myChanel
		fmt.Println("second routune data : ", data)
	}()

	myChanel <- 10
	fmt.Println("End of the main")

}
