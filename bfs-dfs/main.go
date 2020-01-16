package main

import "fmt"

func main() {
	g := make(graph)
	g.Edge(0, 1)
	g.Edge(0, 2)
	g.Edge(1, 2)
	g.Edge(2, 3)
	g.Edge(3, 2)
	g.Edge(3, 4)
	g.Edge(1, 5)

	bfsTree := g.bfs(0, 6)
	fmt.Println(bfsTree)

	dfsTree := g.dfs(0, 6)
	fmt.Println(dfsTree)
}