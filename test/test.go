package main

import (
	"fmt"

	student ".."
	"github.com/01-edu/z01"
)

func print(n int, n2 int) {
	testcase := n
	correct := n2
	if testcase == correct {
		fmt.Print("PASS   ")
	} else {
		fmt.Print("FAIL   ")
	}
	fmt.Printf("%d : %d \n", correct, testcase)
}

func Atoicheck(correct int) int {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"
	n := student.Atoi(s)
	if n == 12345 {
		correct++
	}
	n2 := student.Atoi(s2)
	if n2 == 12345 {
		correct++
	}
	n3 := student.Atoi(s3)
	if n3 == 0 {
		correct++
	}
	n4 := student.Atoi(s4)
	if n4 == 0 {
		correct++
	}
	n5 := student.Atoi(s5)
	if n5 == 1234 {
		correct++
	}
	n6 := student.Atoi(s6)
	if n6 == -1234 {
		correct++
	}
	n7 := student.Atoi(s7)
	if n7 == 0 {
		correct++
	}
	n8 := student.Atoi(s8)
	if n8 == 0 {
		correct++
	}
	return correct
}
func Recursivecheck(correct int) int {
	// case1 : 4^3
	ans1 := student.RecursivePower(4, 3)
	if ans1 == 64 {
		correct++
	}
	// case 2 : 4 ^0
	ans1 = student.RecursivePower(4, 0)
	if ans1 == 1 {
		correct++
	}
	// case 3: 4^-7
	ans1 = student.RecursivePower(4, -7)
	if ans1 == -1 {
		correct++
	}
	return correct
}
func Atoibasecheck(correct int) int {
	n := student.AtoiBase("125", "0123456789")
	if n == 125 {
		correct++
	}
	n = student.AtoiBase("1111101", "01")
	if n == 125 {
		correct++
	}
	n = student.AtoiBase("7D", "0123456789ABCDEF")
	if n == 125 {
		correct++
	}
	n = student.AtoiBase("uoi", "choumi")
	if n == 125 {
		correct++
	}
	n = student.AtoiBase("bbbbbab", "-ab")
	if n == 0 {
		correct++
	}
	return correct

}

// func printnbrbasecheck(correct int) int {
// 	// PrintNbrBase(125, "0123456789")
// 	// z01.PrintRune('\n')
// 	// PrintNbrBase(-125, "01")
// 	// z01.PrintRune('\n')
// 	// PrintNbrBase(125, "0123456789ABCDEF")
// 	// z01.PrintRune('\n')
// 	// PrintNbrBase(-125, "choumi")
// 	// z01.PrintRune('\n')
// 	// PrintNbrBase(125, "aa")
// 	// z01.PrintRune('\n')
// }
// func PrintList(l *student.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *student.NodeI, data int) *student.NodeI {
// 	n := &student.NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }
func main() {
	fmt.Println("Atoi check:")
	print(8, Atoicheck(0))
	fmt.Println("Recursivepower check:")
	print(3, Recursivecheck(0))
	fmt.Println("Printnbrbase check:")
	student.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	student.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	student.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	fmt.Println("Atoibase check:")
	print(5, Atoibasecheck(0))
	// fmt.Println("Splitwhitespace check:")
	// str := "Hello how are you?"
	// fmt.Println(student.SplitWhiteSpaces(str))
	fmt.Println("Advancedsort check:")
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	student.AdvancedSortWordArr(result, student.Compare)
	fmt.Println(result)
	fmt.Println("Activebits check:")
	nbits := student.ActiveBits(7)
	fmt.Println(nbits)
	fmt.Println("Spliwhitespace check:")
	str := "   Hello    how   are you?    "
	fmt.Println(student.SplitWhiteSpaces(str))
	fmt.Println("[Hello how are you?]")

	fmt.Println("Split check:")
	str = "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(str, "Hel"))
	// result := student.ConvertBase("101011", "01", "0123456789")
	// fmt.Println(result)
	// fmt.Println("sortlistinsert check:")

	// var link *student.NodeI

	// link = listPushBack(link, 1)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 9)

	// PrintList(link)

	// link = student.SortListInsert(link, -2)
	// link = student.SortListInsert(link, 2)
	// PrintList(link)

}
