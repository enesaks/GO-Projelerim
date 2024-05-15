package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for index, values := range numbers {
		fmt.Println(index+1, ".Sayı : ", values)
	}
	var languege = "GoLang"

	for _, value := range languege {
		fmt.Printf("%c \n", value)
	}
}
