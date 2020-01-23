package main

import (
	"fmt"
	"os"
)

func stringtonumber(str string) int {
	numb := 0
	check := []rune(str)
	minus := 0
	for ind, s := range check {
		if check[0] == '+' {
			continue
		} else if check[0] == '-' {
			minus++
		} else if s >= '0' && s <= '9' {
			numb *= 10
			numb = int(s) - 48

		} else {
			os.Exit(1)
		}
	}
	if minus > 0 {
		return -numb
	} else {
		return numb
	}

}
func main() {
	count := 0
	for range os.Args {
		count++
	}
	if count > 3 && os.Args[1] == "ztail" && os.Args[2] == "-c" {
		checknum := os.Args[3]
		bytenum := stringtonumber(checknum)
		i := 4
		for i < count {

			name := os.Args[i]
			data, err := os.Open(name)
			if err != nil {
				fmt.Println("open " + name + ": no such file or directory")
				os.Exit(1)

			}
			fmt.Println(name + " :")
			str := os.read(data, bytenum)
			fmt.Println(str)

		}

		os.Exit(0)
	}
	os.Exit(1)
}
