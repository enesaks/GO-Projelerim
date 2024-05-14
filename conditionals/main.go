package main

import "fmt"

func main() {
	var age = 18
	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You can't vote")
	}

	a := 10
	b := 20
	c := 30

	if a >= b && a >= c {
		fmt.Println("a en büyüktür")
	} else if b >= a && b >= c {
		fmt.Println("b en büyüktür")
	} else {
		fmt.Println("c en büyüktür")
	}

}
