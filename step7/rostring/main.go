package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args[1:]) != 1 {
		z01.PrintRune('\n')
		return
	}
	s := []rune(os.Args[1])
	r := []rune{}
	t := true
	for i, char := range s {
		if char != ' ' {
			r = append(r, char)
			s[i] = ' '
			t = false
		} else if !t {
			break
		}
	}
	r2 := []rune{}
	f := true
	for _, char := range s {
		if char != ' ' {
			r2 = append(r2, char)
			f = false
		} else if !f {
			r2 = append(r2, char)
			f = true
		}
	}
	if len(r2) == 0 {
		printer(r)
		z01.PrintRune('\n')
	} else {
		printer(r2)
		z01.PrintRune(' ')
		printer(r)
		z01.PrintRune('\n')
	}
}

func printer(s []rune) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
