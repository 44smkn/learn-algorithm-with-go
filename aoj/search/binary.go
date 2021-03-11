package search

func BinarySearch(a []int, key int) bool {
	left := 0
	right := len(a)

	for left < right {
		mid := (left + right) / 2
		if mid == key {
			return true
		} else if key < a[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}
