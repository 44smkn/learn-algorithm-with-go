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
	counter := 0
	sevenFiveThree(0, n, 0, &counter)
	fmt.Fprintln(w, counter)
}

func sevenFiveThree(current, n, flag int, counter *int) {
	if current > n {
		return
	}
	if flag == 7 {
		*counter += 1
	}

	sevenFiveThree(10*current+3, n, flag|1, counter)
	sevenFiveThree(10*current+5, n, flag|2, counter)
	sevenFiveThree(10*current+7, n, flag|4, counter)
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
