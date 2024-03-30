package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		// os.Stdout.WriteString("\n")
		return
	}
	if os.Args[1] == "" {
		// os.Stdout.WriteString("\n")
		return
	}
	num, _ := strconv.Atoi(os.Args[1])
	if num == 0 {
		// os.Stdout.WriteString("0\n")
		return
	}
	if num < 0 {
		// os.Stdout.WriteString("\n")
		return
	}
	var factors []int
	for i := 2; i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			num /= i
			i--
		}
	}
	for i := 0; i < len(factors); i++ {
		if i == len(factors)-1 {
			fmt.Println(factors[i])
			// os.Stdout.WriteString(strconv.Itoa(factors[i]) + "\n")
		} else {
			fmt.Print(factors[i], "*")
			// os.Stdout.WriteString(strconv.Itoa(factors[i]) + "*")
		}
	}
}
