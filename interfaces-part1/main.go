package main

import "fmt"

type Book struct {
	desi int
}

type Electironic struct {
	desi int
}

func (electironic *Electironic) ShippingCost() int {
	return 10 + (electironic.desi * 3)
}

func (book *Book) ShippingCost() int {
	return 5 + (book.desi * 2)
}

func main() {
	book := &Book{desi: 10}
	book2 := &Book{desi: 20}

	fmt.Println(book.ShippingCost())
	fmt.Println(book2.ShippingCost())

	var books []Book
	books = []Book{Book{desi: 10}, Book{desi: 20}}
	fmt.Println(bookTotalShippingCost(books))

	electironic1 := []Electironic{{desi: 10}, {desi: 20}}
	fmt.Println(bookTotalShippingCost(electironic1)) // #TODO
}

func bookTotalShippingCost(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}
