package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	output := ""

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			output += string((char-'a'+13)%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			output += string((char-'A'+13)%26 + 'A')
		} else {
			output += string(char)
		}
	}

	for _, char := range output {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')

}
