package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func checknumber(str string) bool {
	check := []rune(str)
	for _, x := range check {
		if check[0] == '-' || check[0] == '+' {
			continue
		}
		if x >= '0' && x <= '9' {
			continue
		}
		return false
	}
	return true
}
func checkoperator(str string) bool {
	if str == "+" || str == "-" || str == "*" || str == "%" || str == "/" {
		return true
	}
	return false
}
func stringtonumber(str string) int {
	check := []rune(str)
	ans := 0
	start := 0
	if check[0] == '-' {
		start = 1
	}
	ln := 0
	for range check {
		ln++
	}
	for i := start; i < ln; i++ {
		ans *= 10
		ans = ans + int(check[i]) - 48

	}
	if check[0] == '-' {
		return -ans
	}
	return ans
}
func main() {
	ln := 0
	for range os.Args {
		ln++
	}
	if ln == 4 {
		value := os.Args[1]
		operator := os.Args[2]
		vl2 := os.Args[3]
		if checknumber(value) == true && checknumber(vl2) == true && checkoperator(operator) == true {
			v1 := stringtonumber(value)
			v2 := stringtonumber(vl2)
			if operator == "+" {
				fmt.Println(v1 + v2)
			} else if operator == "*" {
				fmt.Println(v1 * v2)
			} else if operator == "-" {
				fmt.Println(v1 - v2)
			} else if operator == "/" {
				if v2 == 0 {
					fmt.Println("No division by 0")
				} else {
					fmt.Println(v1 / v2)
				}

			} else if operator == "%" {
				if v2 == 0 {
					fmt.Println("No Modulo by 0")
				} else {
					fmt.Println(v1 % v2)
				}
			} else {

			}
		} else {
			z01.PrintRune('0')
			z01.PrintRune('\n')
		}
	}
}
