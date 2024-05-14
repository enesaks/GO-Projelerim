package main

import (
	"fmt"
	"reflect"
)

func main() {

	var productName string
	var quantity int
	var isInStock bool = true
	productName = "golang"
	quantity = 10

	fmt.Println(productName, reflect.TypeOf(productName))
	fmt.Println(quantity, reflect.TypeOf(quantity))
	fmt.Println(isInStock, reflect.TypeOf(isInStock))
	deneme := "sa"
	fmt.Println(deneme)

	fmt.Printf("Profuct Name : %s", productName)
}
