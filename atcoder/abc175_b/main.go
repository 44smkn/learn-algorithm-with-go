package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

// https://atcoder.jp/contests/abc175/editorial/50
func solve(r io.Reader, w io.Writer) {
	scanner := initScanner(r)
	n := scanInt(scanner)[0]
	ls := scanInt(scanner)
	sort.Ints(ls)
	count := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				// 三角形の成立条件：https://mathtrain.jp/seiritu
				if ls[i] != ls[j] && ls[j] != ls[k] && ls[k] < ls[j]+ls[i] {
					count += 1
				}
			}
		}
	}
	fmt.Fprintln(w, count)
}

// 間違えた解き方
func mistake(r io.Reader, w io.Writer) {
	scanner := initScanner(r)
	scanInt64(scanner) // nの入力があるが捨てる
	ls := scanUint64(scanner)
	set := make(map[uint64]bool)
	for _, v := range ls {
		set[v] = true
	}
	fmt.Fprintln(w, combin(uint64(len(set)), 3))
}

// utilityメソッド群

const maxCapacity = 512 * 1024 // デフォルト値は64*1024

func initScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scanUint64(scanner *bufio.Scanner) uint64 {
	scanner.Scan()
	val, err := strconv.ParseUint(scanner.Text(), 10, 64)
	check(err)
	return val
}

func scanInt(scanner *bufio.Scanner) (input []int) {
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), " ") {
		val, err := strconv.Atoi(v)
		check(err)
		input = append(input, val)
	}
	return
}

func scanInt64(scanner *bufio.Scanner) (input []int64) {
	scanner.Scan()
	for _, v := range strings.Split(scanner.Text(), " ") {
		val, err := strconv.ParseInt(v, 10, 64)
		check(err)
		input = append(input, val)
	}
	return
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func combin(n, r uint64) uint64 {
	return factorial(n) / (factorial(r) * factorial(n-r))
}
