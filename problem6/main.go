package main

import (
	"fmt"
)

// Fungsi untuk mengecek apakah sebuah bilangan adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// // Fungsi untuk mengecek apakah digit dalam sebuah bilangan adalah bilangan prima
func isFullPrime(n int) bool {
	if !isPrime(n) {
		return false
	}

	for n > 0 {
		digit := n % 10
		if !isPrime(digit) {
			return false
		}
		n /= 10
	}
	return true
}
func main() {
	var number int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&number)


	fmt.Println(isFullPrime(number))
}
