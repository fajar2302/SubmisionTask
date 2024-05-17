package main

import (
	"fmt"
)
func main() {
	var r float64
	var t float64
	var pi float64 = 3.14

	// Meminta pengguna memasukkan nilai jari-jari dan tinggi tabung
	fmt.Print("Masukkan jari-jari tabung: ")
	fmt.Scanln(&r)

	fmt.Print("Masukkan tinggi tabung: ")
	fmt.Scanln(&t)

	// Menghitung luas permukaan tabung
	luasPermukaan := 2 * pi * r * (r + t)

	// Menampilkan hasil
	fmt.Println("Luas permukaan tabung adalah:", luasPermukaan)

}