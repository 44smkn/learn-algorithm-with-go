package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	s := scanRowText()
	t := scanRowText()

	min := len(t)
	for i := 0; i <= len(s)-len(t); i++ {
		substr := s[i : len(t)+i]

		replaceChar := len(t)
		for j := 0; j < len(substr); j++ {
			if substr[j] == t[j] {
				replaceChar--
			}
		}
		if min > replaceChar {
			min = replaceChar
		}
	}

	fmt.Println(min)
}

var once sync.Once
var scanner *bufio.Scanner

func scanRowText() string {
	once.Do(func() {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
	})
	scanner.Scan()
	return scanner.Text()
}
