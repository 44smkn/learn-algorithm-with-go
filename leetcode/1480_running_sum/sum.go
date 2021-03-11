package sum

// https://leetcode.com/problems/running-sum-of-1d-array/
func RunningSum(nums []int) []int {
	output := make([]int, 0, len(nums))
	for i := range nums {
		sum := 0
		for j := 0; j <= i; j++ {
			sum += nums[j]
		}
		output = append(output, sum)
	}
	return output
}
