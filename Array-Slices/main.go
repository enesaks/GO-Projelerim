package main

import "fmt"

func main() {
	//ARRAYS
	/*var names [3]string

	names[0] = "enes"
	names[1] = "zeynep"
	names[2] = "ahmet"

	var names = [4]string{"enes", "zeynep", "ahmet", "mehmet"}

	fmt.Println(names[0:2])*/

	//Slices
	var names = []string{"enes", "zeynep", "ahmet"}
	names = append(names, "mehmet")
	fmt.Println(names)
}
