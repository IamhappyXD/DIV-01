package student

import "fmt"

func ln(str string) int {
	l := 0
	for range str {
		l++
	}
	return l
}
func Split(str, charset string) []string {
	occurance := 1
	st := []rune(str)
	ch := []rune(charset)
	i2 := 0
	ch_ln := ln(charset)
	start := 0
	for i := 0; i < ln(str); i++ {
		if st[i] == ch[i2] {
			i2++
			if i2 == ch_ln && i != ln(str)-1 {
				i2 = 0
				if start != 0 {
					occurance++
				}
			}
		} else {
			start++
		}
	}
	fmt.Println(occurance)
	ans := []string{"a", "b", "c", "d"}

	//ans := make([]string, occurance+1)
	// j := 0
	// ans[0] = ""
	// i2 = 0
	// for i := 0; i < ln(str); i++ {
	// 	if st[i] == ch[i2] {
	// 		i2++
	// 		if i2 == ch_ln && i != ln(str)-1 {
	// 			i2 = 0
	// 			if start != 0 {
	// 				j++
	// 			}
	// 		}
	// 	} else {
	// 		ans[j] += string(st[i])
	// 	}
	// }
	return ans
}
