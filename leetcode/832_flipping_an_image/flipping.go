package flipping

import "math"

func FlipAndInvertImage(A [][]int) [][]int {
	result := make([][]int, 0, len(A))
	for _, row := range A {
		cnt := len(row)
		converted := make([]int, cnt)
		for i, v := range row {
			converted[cnt-(i+1)] = int(math.Abs(float64(v - 1)))
		}
		result = append(result, converted)
	}
	return result
}
