package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

// KOD BAŞLANGICI
func HashCode(dec string) string {
	result := ""

	for _, char := range dec {
		result += string(rune(int(rune(char)) + len(dec)))
	}

	return result
}
