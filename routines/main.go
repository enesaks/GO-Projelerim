package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// go func() {
	// 	defer wg.Done()
	// 	f1()
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	f2()
	// }()

	// wg.Wait()
	// fmt.Println("Ana programın sonu")

	startTime := time.Now()

	go func() {
		defer wg.Done()
		f1()
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer wg.Done()
		f2()
		time.Sleep(3 * time.Second)
	}()
	go func() {
		defer wg.Done()
		f3()
		time.Sleep(3 * time.Second)
	}()

	wg.Wait()
	fmt.Println("Passed Time : ", time.Now().Sub(startTime))

}

func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func f3() {
	fmt.Println("f3")
}
