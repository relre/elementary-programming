package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) == 0 {
		return
	}

	word := os.Args[1]
	noVowels := "No vowels"
	vowelIndex := findFirstVowel(word)

	if vowelIndex == -1 {
		for _, char := range noVowels {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')

	} else {
		word = word[vowelIndex:] + word[:vowelIndex] + "ay"
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}

}

func isVowel(char rune) bool {
	switch char {
	case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
		return true
	default:
		return false
	}
}

func findFirstVowel(s string) int {
	for i, char := range s {
		if isVowel(char) {
			return i
		}
	}
	return -1
}
