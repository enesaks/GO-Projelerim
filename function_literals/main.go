package main

import (
	"fmt"
	"reflect"
)

func main() {

	func() {
		fmt.Println("f1")
	}()

	add := func(x int, y int) int {
		return x + y
	}

	sub := func(x int, y int) int {
		return x - y
	}

	fmt.Println(reflect.TypeOf(add))

	fmt.Println(add(3, 5))

	fmt.Println(calculater(5, 3, add, sub))

}

func calculater(x int, y int, f1 func(x int, y int) int, f2 func(x int, y int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
