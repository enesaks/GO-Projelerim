package main

import "fmt"

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (xecoder *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ide encode edildi")
}

func (xecoder *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ide encode edlemedi")
}

func (yecoder *YEncoder) Encode(value string) {
	fmt.Println("YEncoder ide encode edildi")
}

func (yecoder *YEncoder) Decode(value string) {
	fmt.Println("YEncoder ide encode edlemedi")
}

func main() {

	/*var xecoder XEncoder
	xecoder = XEncoder{}
	xecoder.Encode("12344i5i")

	var yecoder YEncoder
	yecoder = YEncoder{}
	yecoder.Encode("12344i5i")*/

	var xecoder IEncoder
	xecoder = &XEncoder{}
	xecoder.Encode("12344i5i")

	var yecoder IEncoder
	yecoder = &YEncoder{}
	yecoder.Encode("12344i5i")

}
