package student

func tostring(nbr int) string {
	n := nbr
	dig := 0
	for n != 0 {
		dig++
	}
	ans := ""
	for j := dig - 1; j >= 0; j-- {
		ans += string((nbr % 10) + 48)
		nbr /= 10
	}
	return ans
}
func ConvertBase(nbr, baseFrom, baseTo string) string {
	firstnumber := AtoiBase(nbr, baseFrom)
	final := tostring(AtoiBase(tostring(firstnumber), baseTo))
	return final

}
