package divideconquer

func ExhaustiveSearch(A, m []int) (result string) {
	for i := 0; i < len(m); i++ {
		if solve(0, m[i], A) {
			result += "yes"
		} else {
			result += "no"
		}
	}
	return
}

// solve はより小さい部分問題である solve(i+1, m) と solve(i+1, m-A[i])に分割できる
// m: Aを足し合わせて作れるか試すターゲットの数
// i: Aのスライスを走査するインデックス
// 例: m[0]=2, i=3 -> A[4] == 2 であるかという問題 と A[4] - A[0 or 1 or 2 or 3] == 2 であるかという問題
// ※ or の部分は再帰で解いていく
func solve(i, m int, A []int) bool {
	if m == 0 {
		return true
	}
	if i >= len(A) {
		return false
	}
	res := solve(i+1, m, A) || solve(i+1, m-A[i], A)
	return res
}
