package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(reverse(n))
	fmt.Println(reverse(n))

	rot := 4
	fmt.Printf("Rotate %d to the left: %+v\n", rot, rotateLeft(n, rot))
}

func reverse(n []int) []int {
	// Start:
	// i = 0, j = length of n - 1
	// Continue:
	// continue looping while i is less than j
	// Increment:
	// add 1 to i for every loop
	// minus 1 from j for every loop
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		// Reverse:
		// the first element will have the last element value
		// the last element wil have the first element value
		n[i], n[j] = n[j], n[i]
	}
	return n
}

func rotateLeft(num []int, r int) []int {
	var res []int
	first := num[:r]
	last := num[r:]
	res = append(last, first...)
	return res
}
