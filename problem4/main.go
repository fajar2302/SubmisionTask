package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t string

	fmt.Print("Masukkan Kata yang ingin di Enkripsi: ")
	scanner := bufio.NewScanner(os.Stdin)
 	scanner.Scan()
	t = scanner.Text()
	enkripsi(t)
	
}

func enkripsi(text string) {
	lower := strings.ToLower(text)
	jumlah := len((lower))
	for i:= 0; i < jumlah; i++ {
		if lower[i] >= 97 && lower[i] <= 122 {
			output := lower[i]+10
			if output > 122 {
				output -= 26
			}
			fmt.Print(string(output))
		} else {
			fmt.Print(string(lower[i]))
		}
	}
}