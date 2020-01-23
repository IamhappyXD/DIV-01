package student

func strlen(s string) int {
	check := []rune(s)
	ln := 0
	for range check {
		ln++
	}
	return ln
}
func number(s string) bool {
	check := []rune(s)
	i := 0
	ln := strlen(s)
	// Case 1: +12 and -12
	if check[0] == '+' || check[0] == '-' {
		i = 1
	}
	for i != ln {
		if check[i] < '0' || check[i] > '9' {
			return false
		}
		i++
	}
	return true
}
func Atoi(s string) int {
	if number(s) == true {
		check := []rune(s)
		i := 0
		ln := strlen(s)
		numb := 0
		sign := 1
		if check[0] == '+' || check[0] == '-' {
			if check[0] == '-' {
				sign = -1
			}
			i = 1
		}

		for i != ln {
			numb *= 10
			for ch := '0'; ch != check[i]; ch++ {
				numb++
			}
			i++
		}
		return numb * sign
	} else {
		return 0
	}

}
