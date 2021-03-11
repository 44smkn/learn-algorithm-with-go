package shuffle

func Shuffle(nums []int, n int) []int {
	x := nums[:n]
	y := nums[n:]
	output := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		output = append(output, x[i])
		output = append(output, y[i])
	}
	return output
}
