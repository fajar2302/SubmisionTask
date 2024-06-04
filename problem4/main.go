package main

import "fmt"

func binarySearch(arr []int, x int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == x {
			return mid // Nilai x ditemukan di indeks mid
		} else if arr[mid] < x {
			left = mid + 1 // Cari di separuh kanan
		} else {
			right = mid - 1 // Cari di separuh kiri
		}
	}

	return -1 // x tidak ditemukan dalam array
}

func main() {
	arr1 := []int{1, 1, 3, 5, 5, 6, 7}
	x1 := 3
	fmt.Println(binarySearch(arr1, x1)) // Output: 2

	arr2 := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	x2 := 100
	fmt.Println(binarySearch(arr2, x2)) // Output: -1
	
	arr3 := []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	x3 := 53
	fmt.Println(binarySearch(arr3, x3))
}
