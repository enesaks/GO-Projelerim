package main

import (
	"fmt"
)

func main() {
	index := 0

	for index <= 3 {
		fmt.Println(index)
		index++
	}
	total := 0
	for i := 1; i <= 3; i++ {
		total += i
		fmt.Println(total)
	}
	index = 0
	var names = [3]string{"enes", "zeynep", "ahmet"}

	for index < len(names) {
		fmt.Println(names[index])
		index++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 3 {
			fmt.Println("3 oldu")
			break
		}
	}

	for i := 0; i < 5; i++ {
		if i == 1 || i == 3 {
			fmt.Println("1 ile 3 basılmayacak ! ")
			continue
		}
		fmt.Println(i)
	}

}
