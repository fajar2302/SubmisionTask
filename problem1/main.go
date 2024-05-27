package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", compare("AKA", "AKASHI"))
}

func compare(kata1, kata2 string) string {
	position := make(map[rune]int)
	for i, char := range kata1 {
		position[char] = i
	}

	var commonLetters []rune
	for _, char := range kata2 {
		if _, exists := position[char]; exists {
			commonLetters = append(commonLetters, char)
		}
	}

	return string(commonLetters)
}