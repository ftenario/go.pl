package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	str := "Hello! How can I help you? ... \nIf you don't say a word, I cant help you!"

	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	var cnt int
	for scanner.Scan() {
		cnt++
	}
	fmt.Printf("No of words: %d\n", cnt)

	scanner = bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanLines)
	var lines int
	for scanner.Scan() {
		lines++
	}
	fmt.Printf("No of lines: %d\n", lines)
}
