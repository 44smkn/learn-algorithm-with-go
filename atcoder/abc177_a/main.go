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
	distance := scanSplitedFloat64()
	time := scanSplitedFloat64()
	speed := scanSplitedFloat64()

	if distance/speed > time {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}

var once sync.Once
var scanner *bufio.Scanner

func scanSplitedFloat64() float64 {
	once.Do(func() {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
	})
	scanner.Scan()
	val, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	return val
}
