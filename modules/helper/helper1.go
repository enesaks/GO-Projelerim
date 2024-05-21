package helper // helperın altında olduğu için paketi farklı.

import "fmt"

func Helper1() { // İsmini buyuk harfle tanımlarsa diger paketlerden erişilebilir.
	fmt.Println("Helper paketi içerisinden çalıştı!!")
	Helper2()
	fmt.Println(helper2değişkeni)
	// aynı paketdeki sınıflar fonksiyonlarına veya değişkenlerine ulaşabilirler.
}
func helper2() { // kucuk harfle başlarsak ulaşılamaz farklı paket içerisinden!

}
