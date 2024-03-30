package main

import "fmt"

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

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
