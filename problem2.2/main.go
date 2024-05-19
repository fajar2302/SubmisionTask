package main

import "fmt"

func main() {
	var angka int
	fmt.Println("Masukkan bilangan: ")
	fmt.Scan(&angka)

	fmt.Println("Faktor dari bilangan",angka,"adalah")

	for i := angka; i >= 1; i-- {
		if angka % i == 0 {
			fmt.Println(i)
		}
	}
}