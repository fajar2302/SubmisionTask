package main

import "fmt"

func main() {
	var harga, diskon float64 

	fmt.Print("Masukkan harga: ")
	fmt.Scanln(&harga)

	fmt.Println("Masukkan diskon (dalam %): ")
	fmt.Scanln(&diskon)

	diskon = diskon / 100
	jmlhDiskon := harga * diskon
	hargastlhDiskon := harga - jmlhDiskon

	fmt.Println("Harga Setelah Diskon: ", hargastlhDiskon)
}
