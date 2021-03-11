/*
	Package search は「探索」のアルゴリズム群です。
	キーと値が同一の単純なデータを探索する問題を解いていく。

	Liner Search

	goの言語レベルでサポートされている
		for i, v range slice { ... }
	と番兵を利用した
		slice = append(slice, key)
		for slive[i] != key { i++ }
	をベンチマークで比較すると
		pkg: aoj/search
		BenchmarkSearchBySentinel-4   	1000000000	         0.0684 ns/op	       0 B/op	       0 allocs/op
		BenchmarkSearchByRange-4      	1000000000	         0.0209 ns/op	       0 B/op	       0 allocs/op
	のような結果となる
*/
package search
