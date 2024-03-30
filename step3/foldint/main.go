package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()
	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
func Add(a, b int) int {
	return a + b
}
func Mul(a, b int) int {
	return a * b
}
func Sub(a, b int) int {
	return a - b
}

// KOD BAŞLANGICI
func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for _, v := range a {
		result = f(result, v)
	}
	for _, char := range Itoa(result) {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
func Itoa(n int) string {
	var s string
	negative := false
	if n == 0 {
		return "0"
	}
	if n < 0 {
		negative = true
		n = -n
	}
	for n > 0 {
		s = string(rune(n%10+'0')) + s
		n /= 10
	}
	if negative {
		s = "-" + s
	}
	return s
}
