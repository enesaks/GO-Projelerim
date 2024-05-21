package main // aynı dizindeki dosyaların paketi aynı olmalıdır

import (
	"fmt"
	"myModule/modules/helper"
)

func main() {
	fmt.Println("Hello")

	helper.Helper1()
}
