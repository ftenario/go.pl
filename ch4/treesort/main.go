package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, val int) *tree {
	//if the *tree is nil, lets create a new tree
	fmt.Println("t:", t)
	if t == nil {
		t = new(tree)
		t.value = val
		return t
	}

	if val < t.value {
		t.left = add(t.left, val)
	} else if val > t.value {
		t.right = add(t.right, val)
	}

	return t
}

func main() {

	Sort([]int{5, 1, 2, 4, 3})
	// fmt.Println(t)
}
