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
	m := scanner.int()
	k := scanner.uint64()

	a := make([]uint64, 1, n) // 先頭に0のレコードを入れる
	var asum uint64
	for i := 0; i < n; i++ {
		asum += scanner.uint64()
		a = append(a, asum)
	}
	b := make([]uint64, 1, m) // 先頭に0のレコードを入れる
	var bsum uint64
	for i := 0; i < m; i++ {
		bsum += scanner.uint64()
		b = append(b, bsum)
	}

	ans := 0
	best := m

	for i := 0; i <= n; i++ {
		if a[i] > k {
			break
		}

		for j := best; j >= 0; j-- {
			if a[i]+b[j] <= k {
				best = j
				if i+j > ans {
					ans = i + j
				}
				break
			}
		}
	}
	// 解く
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

func (s *Scanner) uint64() uint64 {
	s.scanner.Scan()
	val, err := strconv.ParseUint(s.scanner.Text(), 10, 64)
	check(err)
	return val
}

// 64bit端末ならint64と扱われるため
func (s *Scanner) int() int {
	s.scanner.Scan()
	val, err := strconv.Atoi(s.scanner.Text())
	check(err)
	return val
}
