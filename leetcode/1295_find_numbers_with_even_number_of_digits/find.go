package findnums

func FindNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		if 10 <= v && v < 100 || 1000 <= v && v < 10000 || v == 100000 {
			count++
		}
	}
	return count
}
