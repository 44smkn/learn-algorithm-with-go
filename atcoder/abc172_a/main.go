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
	a := scanner.float64()
	ans := int(a + math.Pow(a, 2) + math.Pow(a, 3))
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

func (s *Scanner) float64() float64 {
	s.scanner.Scan()
	val, err := strconv.ParseFloat(s.scanner.Text(), 64)
	check(err)
	return val
}
