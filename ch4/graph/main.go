package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("foo", "bar")
	addEdge("one", "1")

	fmt.Printf("Is %s in the graph? %b\n", "foo", hasEdge("foo", "bar"))
	fmt.Println(graph)
}
