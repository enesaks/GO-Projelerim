package main

import "fmt"

func main() {
	myChanel := make(chan string)

	go func() {
		myChanel <- "hello world!"
	}()

	message, isOpen := <-myChanel

	fmt.Println(message, isOpen)
}
