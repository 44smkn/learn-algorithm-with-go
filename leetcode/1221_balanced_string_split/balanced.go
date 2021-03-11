package balanced

// https://leetcode.com/problems/split-a-string-in-balanced-strings/
func BalancedStringSplit(s string) int {
	var res, cnt int
	for _, c := range s {
		if c == 'L' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			res++
		}
	}
	return res
}
