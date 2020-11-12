package main

import "fmt"

func sum(vals ...int) int {
	var total int
	for _, v := range vals {
		total = total + v
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(9, 5, 1))
	fmt.Println(sum(1, 2, 3, 4, 5))

	val := []int{1, 2, 3, 4}
	fmt.Println(sum(val...))
}
