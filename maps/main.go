package main

import "fmt"

func main() {
	var names map[string]int
	names = make(map[string]int, 0)
	names["enes"] = 1
	names["zeynep"] = 2
	names["Ahmet"] = 3

	fmt.Println(names)

	namess := map[string]int{
		"enes":   1,
		"zeynep": 2,
		"ahmet":  3,
	}
	delete(namess, "ahmet")

	fmt.Println(namess)
}
