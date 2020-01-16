package main

import "container/list"

type set map[int]bool

func (s set) Add(u int) {
	s[u] = true
}

func (s set) Remove(u int) {
	delete(s, u)
}

// graph datastructure using set
type graph map[int]set

func (g graph) Edge(u, v int) {
	if g[u] == nil {
		g[u] = make(set)
	}
	g[u].Add(v)
}

func (g graph) RemoveEdge(u, v int) {
	g[u].Remove(v)
}

// import "container/list"
type queue struct {
	*list.List
}

func (q *queue) Push(n int) {
	if q.List == nil {
		q.List = list.New()
	}
	q.PushBack(n)
}

func (q *queue) Pop() int {
	val := q.Remove(q.Front())
	return val.(int)
}

func (q *queue) Empty() bool {
	return q.Len() == 0
}


// import "container/list"
type stack struct {
	*list.List
}

func (s *stack) Push(n int) {
	if s.List == nil {
		s.List = list.New()
	}
	s.PushBack(n)
}

func (s *stack) Pop() int {
	val := s.Remove(s.Back())
	return val.(int)
}

func (s *stack) Empty() bool {
	return s.Len() == 0
}