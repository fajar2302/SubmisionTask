package main

import "fmt"
func swap(a,b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("nilai a adalah", a)
	fmt.Println("nilai b adalah", b)
}