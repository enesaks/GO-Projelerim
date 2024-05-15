package main

import "fmt"

func main() {
	total := add(2, 3)
	fmt.Println(total)

	total1, fark := calculation(2, 3)
	fmt.Println(total1, fark)

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumNumber := sum(numbers)
	fmt.Println("numbers 'ın toplamı :", sumNumber)

	fmt.Println(sum2(1, 2, 3, 4, 56, 7, 7, 45))
}
func sum2(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func sum(numbers []int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func calculation(x int, y int) (int, int) {
	return x + y, x - y
}

func add(x int, y int) int {
	fmt.Println("Toplam : ", x+y)
	return x + y
}
