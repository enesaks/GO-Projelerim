package main

import (
	"fmt"
)

type ValidationError struct {
	text string
	code int
}

func (validationErorr ValidationError) Error() string {
	return fmt.Sprintf("Hata: %s, Kod : %d", validationErorr.text, validationErorr.code)
}

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) add(product Product) error {
	if len(product.name) == 0 {
		return ValidationError{text: "Ürün ismi boş olamaz!!", code: 1001}
	}
	if (product.price) <= 0 {
		return ValidationError{text: "0 dan büyük olmalıdır!!", code: 1000}
	}

	fmt.Println("Ürün Eklenendi")

	return nil
}

func main() {
	// var pointer1 *int

	// if pointer1 == nil {
	// 	fmt.Println("Pointer Boş!!")
	// }

	// fileData, err := os.ReadFile("text.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(fileData)
	// }

	// result, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	var productService ProductService
	masa := Product{1, "", 1000}
	err := productService.add(masa)

	if err != nil {
		fmt.Println(err)
	}

}

// func divide(x int, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("y, 0 olamaz!")
// 	}

// 	return x / y, nil
// }
