package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isvowel(str rune) bool {
	if str == 'A' || str == 'E' || str == 'I' || str == 'O' || str == 'U' {
		return true
	} else if str == 'a' || str == 'e' || str == 'i' || str == 'o' || str == 'u' {
		return true
	} else {
		return false
	}
}
func main() {
	ln := 0
	for range os.Args {
		ln++
	}
	if ln > 1 {
		newstr := ""
		for i := 1; i != ln; i++ {
			newstr += os.Args[i]
			if i != ln-1 {
				newstr += " "
			}
		}
		ptr := []rune(newstr)
		pln := 0
		for range ptr {
			pln++
		}
		med := pln / 2
		j := pln - 1
		for i := 0; i < med; i++ {
			if isvowel(ptr[i]) {
				for j != med {
					if isvowel(ptr[j]) {
						temp := ptr[i]
						ptr[i] = ptr[j]
						ptr[j] = temp
						j--
						break
					}
					j--
				}
			}
		}
		for _, x := range ptr {
			z01.PrintRune(x)
		}

	}

	z01.PrintRune('\n')
}
