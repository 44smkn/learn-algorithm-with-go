package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	scanner := initScanner(r)
	n := scanner.int()
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		e := scanner.int()
		arr = append(arr, e)
	}

	// フレンドリーさが大きい人から到着したほうが心地よさは大きくなる
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	ans := 0
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, arr[0]) // 一番最初の次は1箇所しか挿入場所がない
	for i := 1; i < n; i++ {
		ans += heap.Pop(h).(int)
		// 挿入できる箇所が2箇所
		heap.Push(h, arr[i])
		heap.Push(h, arr[i])
	}

	fmt.Fprintln(w, ans)
}

// utilityメソッド群

type Scanner struct {
	scanner *bufio.Scanner
}

const maxCapacity = 512 * 1024 // デフォルト値は64*1024

func initScanner(r io.Reader) *Scanner {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords) // スペースごとに読み込む
	return &Scanner{scanner: scanner}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 64bit端末ならint64と扱われるため
func (s *Scanner) int() int {
	s.scanner.Scan()
	val, err := strconv.Atoi(s.scanner.Text())
	check(err)
	return val
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
