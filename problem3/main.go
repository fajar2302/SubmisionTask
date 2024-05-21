package main

import "fmt"

func main() {
	var angka int
	var isPrime bool = true

	fmt.Println("Masukkan angka yang ingin dicek: ")
	fmt.Scan(&angka)
	if angka <= 1 {
		isPrime = false
	} else {
		for i := 2; i < angka; i++ {
			if angka%i == 0 {
				isPrime = false
				break
			}
		}
	}

	fmt.Println(isPrime)
}
