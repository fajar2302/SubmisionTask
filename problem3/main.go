package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)
	perkalian(n)
}

func perkalian(number int) {
	fmt.Println("Tabel Perkalian:")
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}
