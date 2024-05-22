package main

import (
	"context"
	"fmt"
	"time"
)

type Product struct {
	Id   int64
	Name string
}

var productChannel = make(chan Product)

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123") // log takibi için kullanılır.

	f1(ctx)
	// ctx, cencel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cencel()
	// go getProductDetailsFromExternalService(10)

	// select {
	// case productDetail := <-productChannel:
	// 	fmt.Println("Ürün detayları getirildi : ", productDetail)

	// case <-ctx.Done():// Context kullanarak bekleme süresini aşan şeyleri durdurabiliriz.
	// 	fmt.Println("Timeout occured, context canceled")
	// }

	fmt.Println("End of the main")
}

func f1(ctx context.Context) {
	fmt.Println("f1", ctx.Value("correlation-id"))
	f2(ctx)
}
func f2(ctx context.Context) {
	fmt.Println("f2", ctx.Value("correlation-id"))
	f3(ctx)
}
func f3(ctx context.Context) {
	fmt.Println("f3", ctx.Value("correlation-id"))
}

func getProductDetailsFromExternalService(productId int64) {
	time.Sleep(2 * time.Second)

	productChannel <- Product{
		Id:   productId,
		Name: "Camaşır makinesi",
	}

}
