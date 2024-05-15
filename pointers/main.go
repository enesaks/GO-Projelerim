package main

import "fmt"

func main() {
	var a int
	var p *int
	a = 10
	p = &a

	*p = 20 //adresden giderek a nın değerini değiştire biliriz.

	fmt.Println("a'nın değeri : ", a)
	fmt.Println("a'nın ramde tutulduğu yer : ", p)
	fmt.Println("a'nın değeri : ", *p)

	var b int
	b = 10
	fmt.Println("b nin ilk değeri : ", b)
	add12(&b)
	fmt.Println("b nin pointerla değiştirilmiş değeri : ", b)

}

func add12(x *int) {
	*x = 12 + *x
}
