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
	height := scanner.int()
	width := scanner.int()
	k := scanner.int()

	// #のときにtrueとなる二次元sliceを作成
	c := make([][]bool, height)
	for i := 0; i < height; i++ {
		c[i] = make([]bool, width)
		row := scanner.string()
		for j := 0; j < width; j++ {
			c[i][j] = row[j] == '#'
		}
	}

	ans := 0 // 黒がk個になるような赤の行列指定の仕方は何通りか

	// 「どの行を赤でmaskするか」の組み合わせをビット列で表現すると、0..2^n-1 となる
	for maskRow := 0; maskRow < 1<<height; maskRow++ {
		// 外側のループと同じように、どの列を赤でmaskするか」の組み合わせをビット列で表現
		for maskColumn := 0; maskColumn < 1<<width; maskColumn++ {
			black := 0
			// c[i][j] でマス目をそれぞれ黒か確認する
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					// マス目が黒 かつ 赤でmaskされていなければカウントアップ
					if c[i][j] && maskRow>>i&1 == 0 && maskColumn>>j&1 == 0 {
						black += 1
					}
				}
			}
			if black == k {
				ans += 1
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

// 64bit端末ならint64と扱われるため
func (s *Scanner) int() int {
	s.scanner.Scan()
	val, err := strconv.Atoi(s.scanner.Text())
	check(err)
	return val
}

func (s *Scanner) string() string {
	s.scanner.Scan()
	return s.scanner.Text()
}
