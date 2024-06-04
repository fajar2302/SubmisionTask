package main

import (
	"fmt"
)

func main() {
    var A, B, C int
    fmt.Println("Masukkan nilai A, B, dan C:")
    fmt.Scan(&A, &B, &C)

    simpleEquations(A, B, C)
}

func simpleEquations(A, B, C int) {
    for x := 1; x <= A; x++ {
        for y := 1; y <= A; y++ {
            for z := 1; z <= A; z++ {
                if x+y+z == A && x*y*z == B && (x*x)+(y*y)+(z*z) == C {
                    fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
                    return
                }
            }
        }
    }
    fmt.Println("Tidak ada solusi yang ditemukan.")
}
