package student

func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if f(array[i], array[j]) == 1 {
				temp := array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
}
