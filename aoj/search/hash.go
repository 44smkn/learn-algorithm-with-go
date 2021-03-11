package search

// Dictionary はハッシュを利用した辞書配列を表します
type Dictionary struct {
	size int
	data []*int
}

// デフォルトの辞書配列のサイズ
const defaultDictSize = 1046527

// NewDictionaryDefaultSize はデフォルトのサイズでの辞書配列を生成します
func NewDictionaryDefaultSize() Dictionary {
	return NewDictionary(defaultDictSize)
}

// NewDictionary は与えられたサイズで辞書配列を生成します
func NewDictionary(size int) Dictionary {
	return Dictionary{
		size: size,
		data: make([]*int, size),
	}
}

func h1(key, size int) int {
	return key % size
}

// 衝突した場合の第2のハッシュ関数
func h2(key, size int) int {
	return 1 + key%(size-1)
}

func h(key, i, size int) int {
	return (h1(key, size) + i*h2(key, size)) % size
}

// Insert は辞書配列にデータを挿入します
func (d Dictionary) Insert(key int) {
	i := 0
	for {
		j := h(key, i, d.size)
		if d.data[j] == nil {
			d.data[j] = &key
			return
		} else {
			i++
		}
	}
}

// Search は辞書配列からデータを探索します
func (d Dictionary) Search(key int) (int, bool) {
	i := 0
	for {
		j := h(key, i, d.size)
		if *d.data[j] == key {
			return j, true
		} else if d.data[j] == nil || i >= d.size {
			return 0, false
		} else {
			i++
		}
	}
}
