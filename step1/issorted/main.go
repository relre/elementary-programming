package main

import (
	"fmt"
)

func main() {
	a1 := []int{-5, -1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	f := func(a, b int) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

// KOD BAŞLANGICI
func IsSorted(f func(int, int) int, a []int) bool {
	ascendingOrdered := true
	descendingOrdered := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			ascendingOrdered = false
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) > 0 {
			descendingOrdered = false
		}
	}

	return ascendingOrdered || descendingOrdered
}
