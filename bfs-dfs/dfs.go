package main

func (g graph) dfs(start, size int) []int {
	visited := make([]bool, size)
	parent := make([]int, size)

	for i := 0; i < size; i++ {
		parent[i] = i
	}

	var s stack
	s.Push(start)
	for !s.Empty() {
		u := s.Pop()
		if visited[u] {
			continue
		}
		visited[u] = true

		for v := range g[u] {
			if !visited[v] {
				parent[v] = u
				s.Push(v)
			}
		}
	}
	return parent
}