package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan tinggi piramida: ")
	fmt.Scanln(&n)

	for i := 0; i <= n; i++ {
		// Membuat spasi sebelum bintang
		for j := 0; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// Membuat bintang dengan nilai berurutan
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
