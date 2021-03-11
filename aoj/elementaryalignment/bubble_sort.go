package elementaryalignment

func BubbleSort(arr []int) int {
	var count int
	// iは未ソート部分列の先頭を指すループ変数で、配列の先頭から末尾に移動する
	for i := 0; i < len(arr)-1; i++ {
		// jは未ソート部分列の隣り合う要素を比較するためのループ変数で、arrの末尾であるN-1から開始し、i+1まで減少する
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				count++
			}
		}
	}
	return count
}
