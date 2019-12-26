package main

import "fmt"

func ActiveBits(n int) uint {
	var bits uint
	bits = 0
	for n > 0 {
		if n%2 == 1 {
			bits++
		}
		n = n / 2
	}
	return bits
}

func main() {
	nbits := ActiveBits(1)
	fmt.Println(nbits)
}
