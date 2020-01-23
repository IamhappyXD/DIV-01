package student

import "fmt"

func SplitWhiteSpaces(str string) []string {
	check := []rune(str)
	count := 1
	ln := 0
	for range check {
		ln++
	}
	start := 0
	for i := 0; i < ln-1; i++ {
		if start != 0 && (check[i] == ' ' || check[i] == '\n' || check[i] == 11 || check[i] == 9) {
			if check[i+1] == ' ' || check[i+1] == '\n' || check[i+1] == 11 || check[i+1] == 9 {
				start--
				continue
			}
			count++
		} else {
			start++
		}
	}
	ans := make([]string, count)
	j := 0
	ans[0] = ""
	start = 0
	for i := 0; i < ln-1; i++ {
		if start != 0 && (check[i] == ' ' || check[i] == '\n' || check[i] == 11 || check[i] == 9) {
			if check[i+1] == ' ' || check[i+1] == '\n' || check[i+1] == 11 || check[i+1] == 9 {
				start--
				continue
			}
			j++
		} else {
			start++
		}
		ans[j] += string(check[i])
	}
	fmt.Println("hjsh")
	for _, s := range ans {
		fmt.Println(s)
	}
	fmt.Println("hjsh")

	return ans

}
