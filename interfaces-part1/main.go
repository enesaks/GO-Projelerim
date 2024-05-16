package main

import "fmt"

type IShipible interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electironic struct {
	desi int
}

type Flower struct {
	desi int
}

func (electironic *Electironic) ShippingCost() int {
	return 10 + (electironic.desi * 3)
}

func (book *Book) ShippingCost() int {
	return 5 + (book.desi * 2)
}
func (flower *Flower) ShippingCost() int {
	return 5 + (flower.desi * 2)
}

func main() {
	book := &Book{desi: 10}
	book2 := &Book{desi: 20}

	fmt.Println(book.ShippingCost())
	fmt.Println(book2.ShippingCost())

	var books []Book
	books = []Book{{desi: 10}, {desi: 20}}
	//fmt.Println(bookTotalShippingCost(books[]))

	electironic1 := []Electironic{{desi: 10}, {desi: 20}}
	//fmt.Println(bookTotalShippingCost(electironic1[]))

	flower := Flower{desi: 10}

	products := []IShipible{&books[0], &electironic1[0], &flower}

	fmt.Println(bookTotalShippingCost(products))
}

func bookTotalShippingCost(products []IShipible) int {
	total := 0
	for _, value := range products {
		total += value.ShippingCost()
	}
	return total
}

/*
func bookTotalShippingCost(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}

func bookTotalShippingCost(electironics []Electironic) int {
	total := 0
	for _, electironic := range electironics {
		total += electironic.ShippingCost()
	}
	return total
}
*/
