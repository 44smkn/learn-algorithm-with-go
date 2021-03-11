package elementaryalignment

func InsertionSort(arr []int) {
	// iは未ソート部分列の先頭インデックス
	for i := 1; i < len(arr); i++ { // 1番目の要素は既にソート済の想定なので、2番目の要素からスタートする
		// jはソート済み部分列からAを挿入するためのループ変数
		j, key := i-1, arr[i]
		for ; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}

		// why j "+1" ?
		// 条件をジャッジする前にjのデクリメントが行われているため
		arr[j+1] = key
	}
}
