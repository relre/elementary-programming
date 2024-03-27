package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {

		// string ifadeyi rune'e çeviriyoruz rune'ler de unicode olarak çalışıyor
		arg := []rune(os.Args[1])
		for i, ch := range arg {
			if ch >= 'a' && ch <= 'z' {
				// z - a + a = kullanımında z çıktısını alırız. o yüzden ortada ki karakteri değişkene bağlarsak z'den aşağıya doğru iner
				arg[i] = 'z' - ch + 'a'
			} else if ch >= 'A' && ch <= 'Z' {
				arg[i] = 'Z' - ch + 'A'
			}
			z01.PrintRune(arg[i])
		}
	}
	z01.PrintRune('\n')
}
