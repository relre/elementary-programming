package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		for _, char := range isPowerOf(nbr) {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')

	}

}

func isPowerOf(n int) string {
	if n == 0 || n == 1 {
		return "false"
	}
	for n != 1 {
		if n%2 != 0 {
			return "false"
		}
		n /= 2
	}
	return "true"
}
