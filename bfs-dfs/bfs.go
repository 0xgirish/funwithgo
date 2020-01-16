package main

func (g graph) bfs(start, size int) []int {
	visited := make([]bool, size)
	parent := make([]int, size)

	for i := 0; i < size; i++ {
		parent[i] = i
	}

	var q queue
	q.Push(start)
	for !q.Empty() {
		u := q.Pop()
		if visited[u] {
			continue
		}
		visited[u] = true

		for v := range g[u] {
			if !visited[v] {
				parent[v] = u
				q.Push(v)
			}
		}
	}
	return parent
}