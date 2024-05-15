package main

import "fmt"

func main() {
	var customer1 Customer

	customer1 = Customer{id: 1, name: "enes", age: 22, address: Address{"istanbul", "esenler"}}

	fmt.Println(customer1)
	customer1.ToString()
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Address
}

type Address struct {
	city     string
	district string
}

func (customer *Customer) ToString() {
	fmt.Printf("name: %s, age: %d", customer.name, customer.age)
}
