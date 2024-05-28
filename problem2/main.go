package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t string
	var angka int
	
	fmt.Print("Masukkan Kata yang ingin di Enkripsi: ")
	scanner := bufio.NewScanner(os.Stdin)
 	scanner.Scan()
	t = scanner.Text()

	fmt.Print("masukkan offset angka: ")
	fmt.Scanln(&angka)
	enkripsi(angka, t)
	
}

func enkripsi(offset int, text string) {
	lower := strings.ToLower(text)

	slice := []rune(lower)
	fmt.Println(slice)


	for i := 0; i < len(slice); i++ {
		if slice[i] >= 97 && slice[i] <= 122 {
			output := slice[i]+rune(offset%26)
			if output > 122 {
				output -= 26
			}
			fmt.Print(string(output))
		} else {
			fmt.Print(string(lower[i]))
		}
    }

}