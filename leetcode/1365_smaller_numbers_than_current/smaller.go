package smaller

// https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/
func SmallerNumbersThanCurrent(nums []int) []int {

	// constraint: 0 <= nums[i] <= 100
	counts := make([]int, 101)
	for _, num := range nums {
		counts[num]++
	}
	for i := 1; i <= 100; i++ {
		counts[i] += counts[i-1]
	}

	output := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			output[i] = 0
		} else {
			output[i] = counts[num-1]
		}
	}

	return output
}
