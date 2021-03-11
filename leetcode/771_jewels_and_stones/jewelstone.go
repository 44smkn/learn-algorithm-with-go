package jewelstone

// https://leetcode.com/problems/jewels-and-stones/
func NumJewelsInStones(J string, S string) int {
	output := 0
	for _, jewel := range J {
		for _, stone := range S {
			if stone == jewel {
				output++
			}
		}
	}
	return output
}
