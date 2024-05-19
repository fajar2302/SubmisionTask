package main

import (
	"fmt"
	"math"
)

func main() {
	var angka int
	var exponen int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&exponen)

	fmt.Println(pangkat(angka, exponen))
}


func pangkat(angka, exponen int) int {
	return int(math.Pow(float64(angka), float64(exponen)))
}