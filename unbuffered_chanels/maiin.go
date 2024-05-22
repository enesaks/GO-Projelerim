package main

import (
	"fmt"
)

func main1() {
	myChanel := make(chan string)
	done := make(chan bool)

	// go func() {
	// 	myChanel <- "hello world!"
	// }()

	go func() {
		message := <-myChanel
		fmt.Println(message)

		done <- true
	}()

	go func() {
		myChanel <- "hellooo"
	}()

	<-done //diğeri çalışana kadar bloklıyacak.
	fmt.Println("End of the main")
}
