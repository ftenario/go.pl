package main

import "fmt"

func main() {
	str := []string{"one", "", "three", "", "five"}
	fmt.Println(str)
	fmt.Println(nonempty(str))
	fmt.Println(str)

	fmt.Println("Removing adjacent duplicate strings")
	str = []string{"one", "two", "two", "three", "three"}
	fmt.Println(removeAdjucentDuplicate(str))

	fmt.Println("Using Deduplication in one pass")
	fmt.Println(deduplication(str))
}

func nonempty(str []string) []string {
	var x int
	for _, v := range str {
		if v != "" {
			// fmt.Printf("index: %d, value: %s\n", x, v)
			str[x] = v
			x++
		}
	}
	return str[:x]
}

//in-place function that removes adjacent duplicate strings
func removeAdjucentDuplicate(str []string) []string {
	var prev string
	var cnt int
	for _, v := range str {
		if prev != v {
			str[cnt] = v
			cnt++
		}
		prev = v
	}
	return str[:cnt]
}

// deduplication - one pass using map
func deduplication(str []string) []string {
	var out []string
	temp := make(map[string]string)

	for _, v := range str {
		if _, ok := temp[v]; !ok {
			temp[v] = ""
			out = append(out, v)
		}
	}
	return out
}
