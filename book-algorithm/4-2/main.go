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

func solve(r io.Reader, w io.Writer) {
	scanner := initScanner(r)
	n := scanner.int()

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	fmt.Fprintln(w, tribonacci(memo, n-1))
}

func tribonacci(memo []int, n int) int {
	if n == 0 || n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	val := tribonacci(memo, n-1) + tribonacci(memo, n-2) + tribonacci(memo, n-3)
	memo[n] = val
	return val
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
