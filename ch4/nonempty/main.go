package main

import "fmt"

func main() {
	str := []string{"one", "", "three", "", "five"}
	fmt.Println(str)
	fmt.Println(nonempty(str))
	fmt.Println(str)
}

func nonempty(str []string) []string {
	var x int
	for _, v := range str {
		if v != "" {
			fmt.Printf("index: %d, value: %s\n", x, v)
			str[x] = v
			x++
		}
	}
	return str[:x]
}
