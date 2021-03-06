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

	var ac uint64
	var wa uint64
	var tle uint64
	var re uint64

	n := scanInt(scanner)
	for i := 0; i < n; i++ {
		c := scanString(scanner)
		switch c {
		case "AC":
			ac += 1
		case "WA":
			wa += 1
		case "TLE":
			tle += 1
		case "RE":
			re += 1
		}
	}
	fmt.Fprintf(w, "AC x %v\nWA x %v\nTLE x %v\nRE x %v\n", ac, wa, tle, re)
}

// utilityメソッド群

const maxCapacity = 512 * 1024 // デフォルト値は64*1024

func initScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Split(bufio.ScanWords) // スペースごとに読み込む
	return scanner
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 64bit端末ならint64と扱われるため
func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	val, err := strconv.Atoi(scanner.Text())
	check(err)
	return val
}

func scanString(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
