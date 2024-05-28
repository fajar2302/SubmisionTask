package main

import (
	"fmt"
)

func maxSubarraySum(arr []int, k int) int {
	if k > len(arr) {
		return -1 // Jika k lebih besar dari banyak angka dalam array, kembalikan -1
	}

	maxSum := 0
	currentSum := 0

	// Menghitung jumlah awal untuk array pertama dengan ukuran 'k'
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}
	currentSum = maxSum

	// Memproses array berikutnya 
	for i := k; i < len(arr); i++ {
		currentSum = currentSum + arr[i] - arr[i-k]

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	arrInput := []int{2, 3, 4, 1, 5}
	k := 2
	maxSum := maxSubarraySum(arrInput, k)
	fmt.Println("Output:", maxSum) // Output: 9
}
