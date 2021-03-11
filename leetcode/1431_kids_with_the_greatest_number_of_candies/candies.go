package candies

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
func KidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}

	output := make([]bool, 0, len(candies))
	for _, v := range candies {
		output = append(output, max <= v+extraCandies)
	}

	return output
}
