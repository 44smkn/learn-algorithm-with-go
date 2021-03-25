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
	s := scanner.string()
	n := len(s)
	ans := 0
	for bit := 0; bit < 1<<(n-1); bit++ {
		tmp := 0
		for i := 0; i < n-1; i++ {
			tmp *= 10
			val, _ := strconv.Atoi(string(s[i]))
			tmp += val
			if bit&(1<<i) > 0 {
				ans += tmp
				tmp = 0
			}
		}
		// 先程までのループだと最後の桁が評価されないので、最後の桁を評価してあげる必要がある
		tmp *= 10
		val, _ := strconv.Atoi(string(s[n-1]))
		tmp += val
		ans += tmp
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

func (s *Scanner) string() string {
	s.scanner.Scan()
	return s.scanner.Text()
}
