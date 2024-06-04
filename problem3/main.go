package main

import (
	"fmt"
	"sort"
)

func minTotalHeight(dragonHeads []int, knights []int) interface{} {
	// Jika jumlah kepala naga lebih besar dari jumlah ksatria, tidak mungkin untuk memotong semua kepala naga
	if len(dragonHeads) > len(knights) {
		return "knight fall"
	}

	// Mengurutkan diameter kepala naga dan tinggi ksatria secara menaik
	sort.Ints(dragonHeads)
	sort.Ints(knights)

	totalHeight := 0
	knightIndex := 0
	// Melakukan iterasi untuk setiap kepala naga
	for _, head := range dragonHeads {
		// Mencari ksatria dengan tinggi minimal yang bisa memotong kepala naga saat ini
		for knightIndex < len(knights) && knights[knightIndex] < head {
			knightIndex++
		}
		// Jika tidak ada ksatria yang bisa memotong kepala naga saat ini, return "knight fall"
		if knightIndex == len(knights) {
			return "knight fall"
		}
		// Menghitung total tinggi ksatria yang digunakan untuk memotong semua kepala naga
		totalHeight += knights[knightIndex]
		knightIndex++
	}

	return totalHeight
}

func main() {
	// Test cases
	dragonHeads1 := []int{5, 4}
	knights1 := []int{7, 8, 4}
	fmt.Println(minTotalHeight(dragonHeads1, knights1)) // Output: 11
}
