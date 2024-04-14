package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

// KOD BAŞLANGIC YERİ
func Split(s, sep string) []string {
	startIndex := 0
	r := []string{}
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep { // if you find sep inside the string
			if s[startIndex:i] != "" { // don't append empty strings
				r = append(r, s[startIndex:i])
			}
			startIndex = i + len(sep)
			// i+=len(sep) // ADD THIS LINE TO INCLUDE THE "sep"
		} else if i+len(sep) == len(s) { // if len s[i:i+len(sp)]!= sep and i+len(sp)=len(s) (WHIICH MEANS we are at the last iteration) WHIICH MEANS if len s[len(s)-len(sep):len(s))]!= sep WHIICH MEANS if the end of the string doesn't end with a sep
			r = append(r, s[startIndex:]) // in this case append all to the end
		}
	}
	return r
}
