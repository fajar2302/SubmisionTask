package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan jumlah baris/kolom: ")
    fmt.Scanln(&n)
	println(DrawXYZ(n))
}

func DrawXYZ(n int) string {
	var output, hasil string
    // Menggunakan dua loop bersarang untuk mencetak angka dari 1 hingga n * n
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            // Menghitung nilai angka sesuai dengan baris dan kolom
            num := i*n + j + 1
			if num%3 == 0 {
                output = "X"
            } else if num%2 == 0 {
                output = "Z"
            } else {
                output = "Y"
            }
			hasil += output + "\t"
        }
        hasil += "\n"
    }
	return hasil
}
