package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var angka string
	var numberArray []int
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&angka)
	angka = strings.ReplaceAll(angka, ",", "")
	jumlah := len(angka)

	
	// konversi dan memasukkan angka kedalam array
	for i := 0; i < jumlah; i++ {
		digit, err := strconv.Atoi(string(angka[i]))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		numberArray = append(numberArray, digit)
	}

	meanMedian(numberArray)
	mean, median := meanMedian(numberArray)
	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Median: %.2f\n", median)
}


func meanMedian(arrayInput[]int)(float64, float64) {
	// Menghitung Mean
	total := 0
	for _, num := range arrayInput {
		total += num
	}
	mean := float64(total) / float64(len(arrayInput))
	// fmt.Printf("Mean: %.2f\n", mean)

	// Menghitung Median
	sort.Ints(arrayInput)
	var median float64
	if len(arrayInput)%2 == 0 {
		median = float64(arrayInput[len(arrayInput)/2-1]+arrayInput[len(arrayInput)/2]) / 2.0
	} else {
		median = float64(arrayInput[len(arrayInput)/2])
	}
	return mean, median
}