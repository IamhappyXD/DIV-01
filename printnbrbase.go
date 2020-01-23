package student

import (
	"github.com/01-edu/z01"
)

func printnumber(nbr int) {
	x := nbr
	dig := 0
	for x != 0 {
		x /= 10
		dig++
	}
	for dig != 0 {
		n := x / RecursivePower(10, dig-1)
		ch := '0'
		for i := 0; i <= n; i++ {
			ch++
		}
		z01.PrintRune(ch)
		dig--
		x = x % RecursivePower(10, dig-1)
	}
	z01.PrintRune('\n')
}

func PrintNbrBase(nbr int, base string) {
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	if base == "0123456789" {
		print(nbr)
	} else if base == "01" {
		input := nbr
		n := 1

		for n*2 <= input {
			n = n * 2
		}
		for n >= 1 {
			print(input / n)
			input = input % n
			n = n / 2
		}
	} else if base == "0123456789ABCDEF" {
		input := nbr
		n := 1

		for n*16 <= input {
			n = n * 16
		}
		for n >= 1 {
			r := input / n
			if r > 9 {
				ch := 'A'
				for i := 10; i < r; i++ {
					ch++
				}
				z01.PrintRune(ch)
			} else {
				print(r)
			}
			input = input % n
			n = n / 16
		}
	} else {
		z01.PrintRune('h')

	}

}
