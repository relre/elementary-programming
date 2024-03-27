package main

import (
	"fmt"
)

// TEST KISMI
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

// KOD BAŞLANGICI

func CamelToSnakeCase(s string) string {

	result := ""

	if len(s) == 0 || !containOnlyAlphabet(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		if isUpper(rune(s[i])) && !isUpper(rune(s[i+1])) && i+1 < len(s) && i != 0 {
			result += "_"
			result += string(s[i])
			// bu kısmı anlamadım
		} else if !isUpper(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			result += string(s[i])
		} else {
			return s
		}
	}

	return result
}

func isUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func containOnlyAlphabet(s string) bool {
	for _, char := range s {
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}
	}
	return true
}
