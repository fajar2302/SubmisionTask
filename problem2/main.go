package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan jumlah baris/kolom: ")
    fmt.Scanln(&n)
	var output string
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

            fmt.Printf("%s\t", output)
        }
        fmt.Println() // Pindah ke baris berikutnya setelah mencetak semua angka pada baris tertentu
    }
}
