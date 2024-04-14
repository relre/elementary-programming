package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	union := args[0] + args[1]
	result := []rune{}
	if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		z01.PrintRune('\n')
		return
	}
	for _, char := range union { // range over union
		found := false // use found so we don't add similar characters
		for _, n := range result {
			if char == n {
				found = true
				break
			}
		}
		if !found {
			result = append(result, char)
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
