package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	word := os.Args[1]
	firstLetter := os.Args[2]
	secondLetter := os.Args[3]
	output := ""

	for _, char := range word {
		if char == rune(firstLetter[0]) {
			output += string(secondLetter)
		} else {
			output += string(char)
		}

	}

	for _, char := range output {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
