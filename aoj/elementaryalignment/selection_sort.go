package elementaryalignment

func SelectionSort(arr []int) int {
	var count int
	for i := 0; i < len(arr)-1; i++ {
		minj := i
		for j := i + 1; j < len(arr); j++ {
			if arr[minj] > arr[j] {
				minj = j
				count++
			}
		}
		arr[i], arr[minj] = arr[minj], arr[i]
	}
	return count
}
