# learn-algorithm

## 概要

主に競プロの問題を解くことでアルゴリズムへの理解を深めていくリポジトリ  
各ディレクトリは下記のような分け方

* aoj: [プログラミングコンテスト攻略のためのアルゴリズムとデータ構造](プログラミングコンテスト攻略のためのアルゴリズムとデータ構造)
* atcoder: [AtCoder](https://atcoder.jp/contests/)
* automaton: [計算理論の基礎 1.オートマトンと言語](https://www.kyoritsu-pub.co.jp/bookdetail/9784320122079)
* leetcode: [leetcode.com](https://leetcode.com/)

## Goの共通テンプレート

```go
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
    // 解く
    fmt.Fprintln(w /* 答え */)
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

func (s *Scanner) float64() float64 {
    s.scanner.Scan()
    val, err := strconv.ParseFloat(s.scanner.Text(), 64)
    check(err)
    return val
}

func (s *Scanner) string() string {
    s.scanner.Scan()
    return s.scanner.Text()
}

// 階乗
func factorial(n uint64) uint64 {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

// 順列
func perm(n, r uint64) uint64 {
    return factorial(n) / factorial(n-r)
}

// 組み合わせ
func combin(n, r uint64) uint64 {
    return factorial(n) / (factorial(r) * factorial(n-r))
}

```
