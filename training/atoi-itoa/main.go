package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Atoi(s string) int {
	num := 0
	minus := 1
	for _, char := range s {
		if char == '-' {
			minus = -1
		} else {
			num += minus * int(char-'0')
			num *= num
		}

	}

	return num

}

func Itoa(n int) string {

	return "x"

}
