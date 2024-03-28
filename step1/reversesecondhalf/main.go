package main

import (
	"fmt"
)

func main() {
	fmt.Print(ReverseSecondHalf("This is the 1st half This is the 2nd half This is the 3nd half"))
	fmt.Print(ReverseSecondHalf(""))
	fmt.Print(ReverseSecondHalf("12345678.."))
}

// KOD BAŞLANGIÇ KISMI

func ReverseSecondHalf(str string) string {
	if len(str) == 0 {
		return "Invalid Output\n"
	} else if len(str) == 1 {
		return str + "\n"
	} else {
		res := ""
		for i := len(str) - 1; i >= int(len(str)/2); i-- {
			res += string(str[i])
		}
		res += "\n"
		return res
	}

}
