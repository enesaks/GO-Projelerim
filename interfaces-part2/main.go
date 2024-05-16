package main

import "fmt"

type XEncoder struct {
}

func (xecoder *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ide encode edildi")
}

func (xecoder *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ide encode edlemedi")
}

func main() {
	var xecoder XEncoder
	xecoder = XEncoder{}
	xecoder.Encode("12344i5i")

}
