package search_test

import (
	"fmt"
	"impl-using-go/aoj/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    []int
		t    []int
		want int
	}{
		{"case1", []int{1, 2, 3, 4, 5}, []int{3, 4, 1}, 3},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := multipleBinarySearch(tt.s, tt.t)
			if tt.want != got {
				t.Errorf("want: %v  got: %v", tt.want, got)
			}
		})
	}
}

func multipleBinarySearch(s []int, t []int) int {
	var count int
	for _, v := range t {
		if exists := search.BinarySearch(s, v); exists {
			count++
		}
	}
	return count
}

func initVar() (a, q []int, n int) {
	a = makeRange(0, 10000000)
	q = []int{3890000, 4900005, 12890000, 9808009}
	n = len(a)
	return
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// 番兵を利用したケース
// 探しているKeyを一番うしろに挿入することで終了条件の比較演算をへらす
// だけど range を使ったほうが早い
func BenchmarkSearchBySentinel(b *testing.B) {
	a, q, n := initVar()

	b.ResetTimer()
	var sum int
	for _, e := range q {
		if search.SearchBySentinel(a, e, n) {
			sum++
		}
	}
}

func BenchmarkSearchByRange(b *testing.B) {
	a, q, n := initVar()

	b.ResetTimer()
	var sum int
	for _, e := range q {
		if search.SearchByRange(a, e, n) {
			sum++
		}
	}
}

func ExampleNewDictionary() {
	a := search.NewDictionary(1234567)
	a.Insert(2245)
	b, ok := a.Search(2245)
	fmt.Println(b, ok)
}
