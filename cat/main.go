package main

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printmessage(name string) {
	str := "open "
	for _, i := range str {
		z01.PrintRune(i)
	}
	for _, i := range name {
		z01.PrintRune(i)
	}
	z01.PrintRune(':')
	str = " no such file or directory\n"
	for _, i := range str {
		z01.PrintRune(i)
	}
}
func main() {
	count := 0
	for range os.Args {
		count++
	}
	if count == 1 {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text += "^C\n"
		for _, ch := range text {
			z01.PrintRune(ch)
		}
	} else if count == 2 {
		name := os.Args[1]
		data, err := ioutil.ReadFile(name)
		if err != nil {
			printmessage(name)
			return
		}
		str := string(data)
		ptr := []rune(str)
		for _, i := range ptr {
			z01.PrintRune(i)
		}
		z01.PrintRune(10)
	} else if count == 3 {
		num := 1
		for num < 3 {
			name := os.Args[num]
			data, err := ioutil.ReadFile(name)
			if err != nil {
				printmessage(name)
				return
			}
			str := string(data)
			ptr := []rune(str)
			for _, i := range ptr {
				z01.PrintRune(i)
			}
			z01.PrintRune('\n')

			num++
		}
	} else {
		return
	}
}
