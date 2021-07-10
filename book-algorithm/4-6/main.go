package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, writer io.Writer) {
	scanner := initScanner(r)
	n := scanner.int()
	w := scanner.int()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanner.int()
	}
	memo := make([][]*bool, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]*bool, w+1)
	}
	result := partialSum(n, w, a, memo)
	fmt.Fprintln(writer, result)
}

func partialSum(i, w int, a []int, memo [][]*bool) bool {
	if i == 0 {
		return w == 0
	}

	if memo[i][w] != nil {
		return *memo[i][w]
	}

	if partialSum(i-1, w-a[i-1], a, memo) {
		memo[i][w] = Bool(true)
		return true
	}

	if partialSum(i-1, w, a, memo) {
		memo[i][w] = Bool(true)
		return true
	}

	memo[i][w] = Bool(false)
	return false
}

func Bool(val bool) *bool {
	b := val
	return &b
}

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
