package elementaryalignment

func ShellSort(arr, gs []int) {
	for _, g := range gs {
		insertionSort(arr, g)
	}
}

func insertionSort(arr []int, g int) {
	for i := g; i < len(arr); i++ { // iは未ソート部分列の先頭(先頭gの数の要素は整列済みとする)
		j := i - g // jはソート済部分列
		key := arr[i]
		for ; j >= 0 && arr[j] > key; j -= g {
			arr[j+g] = arr[j]
		}
		arr[j+g] = key
	}
}
