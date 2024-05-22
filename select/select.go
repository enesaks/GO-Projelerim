package main

import (
	"fmt"
	"time"
)

func main() {

	chanel1 := make(chan string)
	chanel2 := make(chan string)

	var data1 string
	var data2 string

	go func() {
		time.Sleep(5 * time.Second)
		chanel1 <- "helloo"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		chanel2 <- "world "
	}()

	for len(data1) == 0 || len(data2) == 0 {
		select {
		case data1 = <-chanel1:
			fmt.Println("data1: ", data1)
		case data2 = <-chanel2:
			fmt.Println("data2: ", data2)
		default:
			fmt.Println("No data came")
		}
		time.Sleep(1 * time.Second)
	}
}
