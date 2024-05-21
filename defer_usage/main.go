package main

import (
	"fmt"
)

func main() {
	// defer fmt.Println("Fonksiiyon bitince çalışıcak")
	// fmt.Println("hello")
	// defer fmt.Println("ilk defer son çalışır") // oyuzden bu daha önce çalştı.
	// var condition = true
	// defer fmt.Println("1.defer")
	// if condition {
	// 	return
	// }
	// fmt.Println("2.defer")

	// var x = 10
	// var y = 20

	// defer fmt.Println(x)
	// defer fmt.Println(y)

	// x = 100
	// y = 200

	// defer fmt.Println(x)
	// defer fmt.Println(y)

	var condition = true
	defer cleanup()
	if condition {
		panic("An erorr occurred")
	}

}

func cleanup() {
	fmt.Println("Cleanup Worked...")
}
