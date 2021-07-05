package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	scanner := initScanner(r)
	n := scanner.int()
	dc := countDigit(n)
	count := 0
	for i := dc; i >= 3; i-- {
		count += sevenFiveThree(i, 0, n, 0)
	}
	fmt.Fprintln(w, count)
}

func countDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + countDigit(n/10)
}

func sevenFiveThree(digit, sft, n, flag int) int {
	if digit == 0 {
		if sft <= n && flag == 7 {
			return 1
		} else {
			return 0
		}
	}

	return sevenFiveThree(digit-1, sft+3*int(math.Pow10(digit-1)), n, flag|1) + sevenFiveThree(digit-1, sft+5*int(math.Pow10(digit-1)), n, flag|2) + sevenFiveThree(digit-1, sft+7*int(math.Pow10(digit-1)), n, flag|4)
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
