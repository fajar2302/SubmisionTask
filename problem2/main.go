package main

import "fmt"
func main() {
	var alas float64
	var tinggi float64

	fmt.Print("Alas: ")
	fmt.Scanln(&alas)

	fmt.Print("Tinggi: ")
	fmt.Scanln(&tinggi)

	luas := 0.5 * alas * tinggi

	fmt.Println("Luas segitiga adalah:", luas)
}