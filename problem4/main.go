package main

import (
	"fmt"
	"strings"
)

func palindrome(kata string) bool {
	// Menghapus spasi dan mengubah huruf kecil semua
	kata = strings.ToLower(strings.ReplaceAll(kata, " ", ""))

	// Memeriksa apakah kata merupakan palindrom atau bukan
	for i, j := 0, len(kata)-1; i < j; i, j = i+1, j-1 {
		if kata[i] != kata[j] {
			return false
		}
	}
	return true
}

func main() {
	var kata string

	fmt.Println("Masukkan kata yang ingin dicek: ")
	fmt.Scan(&kata)

	fmt.Println(palindrome(kata))
}
