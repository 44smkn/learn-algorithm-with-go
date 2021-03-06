package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	// wrong()
	correct()
}

func correct() {
	n := scanSplitedInt()
	array := make([]uint64, 0, n)

	// 配列を作る
	for i := 0; i < n; i++ {
		array = append(array, scanSplitedUint64())
	}

	// 累積和を作る
	cusum := make([]uint64, n+1)
	for i := 0; i < n; i++ {
		cusum[i+1] = (cusum[i] + array[i])
	}

	var ans uint64
	for i, v := range array {
		// modは計算途中で何回いつ挟んでも良い
		// 参考: https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a
		sum := (cusum[n] - cusum[i+1]) % 1000000007
		ans += v * sum
		ans %= 1000000007
	}

	fmt.Println(ans)
}

// 間違った回答
func wrong() {
	n := scanSplitedInt()
	array := make([]uint64, 0, n)
	for i := 0; i < n; i++ {
		array = append(array, scanSplitedUint64())
	}
	var sum uint64
	for i, v := range array {
		for j := i + 1; j < n; j++ {
			sum += v * array[j]
		}
	}

	fmt.Println(sum % 1000000007)
}

var once sync.Once
var scanner *bufio.Scanner

func scan() {
	once.Do(func() {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
	})
	scanner.Scan()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scanSplitedUint64() uint64 {
	scan()
	val, err := strconv.ParseUint(scanner.Text(), 10, 64)
	check(err)
	return val
}

func scanSplitedInt() int {
	scan()
	val, err := strconv.Atoi(scanner.Text())
	check(err)
	return val
}
