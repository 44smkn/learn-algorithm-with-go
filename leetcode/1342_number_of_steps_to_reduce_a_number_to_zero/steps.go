package steps

func NumberOfSteps(num int) int {
	op := 0
	for {
		if num == 0 {
			return op
		}

		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		op++
	}
}
