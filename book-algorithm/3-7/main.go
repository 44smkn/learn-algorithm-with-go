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
	ans := 0
	for bit := 0; bit < 1<<(len(s)-1); bit++ {
		tmp := 0
		for i := 0; i < len(s); i++ {
			tmp *= 10
			val, _ := strconv.Atoi(string(s[i]))
			tmp += val
			if bit&(1<<i) == 0 {
				ans += tmp
				tmp = 0
			}
		}
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
