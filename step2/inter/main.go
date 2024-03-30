package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	s1 := []rune(rmvRepStr(args[0]))
	s2 := []rune(args[1])
	r := []rune{}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s2[j] == s1[i] {
				r = append(r, s1[i])
				break
			}
		}
	}
	for _, char := range r {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func rmvRepStr(s string) string {
	r := []rune{}
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		rep := false
		for j := 0; j < len(r); j++ {
			if sliceS[i] == r[j] {
				rep = true
				continue
			}
		}
		if !rep {
			r = append(r, sliceS[i])
		}
	}
	return string(r)
}
