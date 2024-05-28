package main

import "fmt"

func main() {
    input1 := []int{1, 2, 3, 4}
    input2 := []int{1, 3, 5, 10, 16}
    
    hasil := PerbedaanArray(input1, input2)
    
    fmt.Println(hasil)
}

func PerbedaanArray(input1, input2 []int) []int {
    countMap := make(map[int]bool)
    for _, num := range input2 {
        countMap[num] = true
    }
    
    var hasil []int
    for _, num := range input1 {
        if !countMap[num] {
            hasil = append(hasil, num)
        }
    }
    
    return hasil
}
