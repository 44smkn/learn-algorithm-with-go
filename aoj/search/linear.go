package search

func SearchByRange(a []int, key, n int) bool {
	for _, v := range a {
		if v == key {
			return true
		}
	}
	return false
}

func SearchBySentinel(a []int, key, n int) bool {
	a = append(a, key)

	// 番兵を利用して比較演算子をへらす
	i := 0
	for a[i] != key {
		i++
	}
	return i != n
}
