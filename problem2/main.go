package main

import (
	"fmt"
)

var denominations = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}

func main() {
    var amount int
    fmt.Println("Masukkan jumlah uang yang ingin ditukar:")
    fmt.Scan(&amount)

    change := makeChange(amount)
    fmt.Println("Pecahan uang yang diberikan:", change)
}

func makeChange(amount int) []int {
    change := make([]int, 0)

    for _, denomination := range denominations {
        if amount >= denomination {
            count := amount / denomination
            for i := 0; i < count; i++ {
                change = append(change, denomination)
            }
            amount -= count * denomination
        }
    }

    return change
}
