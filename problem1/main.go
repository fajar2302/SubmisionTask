package main

import "fmt"

func main() {
	var nilai float32
	var predikat string
	fmt.Println("Masukkan Nilai: ")
	fmt.Scanln(&nilai)

	if (nilai >= 80 && nilai <= 100){
		predikat = "A"
	} else if (nilai >= 65 && nilai < 80 ) {
		predikat = "B+"	
	} else if (nilai >= 50 && nilai < 65) {
		predikat = "B"
	} else if (nilai >= 35 && nilai < 50) {
		predikat = "C"
	} else if (nilai >= 0 && nilai < 35) {
		predikat = "D"
	}

	fmt.Println("Predikatnnya: ", predikat)

}