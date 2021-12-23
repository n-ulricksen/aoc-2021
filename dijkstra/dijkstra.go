package dijkstra

import (
	"container/heap"
	"math"
)

// Graph node
type Node struct {
	Weight   int
	Priority int
}

// Minimum Priority Queue
type minPQ []*Node

func (pq minPQ) Len() int { return len(pq) }

func (pq minPQ) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq minPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *minPQ) Push(x interface{}) {
	item := x.(*Node)
	(*pq) = append(*pq, item)
}

func (pq *minPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	ret := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[:n-1]
	return ret
}

func Dijkstra(graph map[*Node][]*Node, start *Node,
	end *Node) map[*Node]int {

	// Initialize distance map
	dist := make(map[*Node]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	// Initialize queue
	queue := &minPQ{}
	heap.Init(queue)
	heap.Push(queue, start)

	visited := make(map[*Node]bool)

	// Compute!
	for queue.Len() > 0 {
		item := heap.Pop(queue).(*Node)
		visited[item] = true

		for _, neighbor := range graph[item] {
			if visited[neighbor] {
				continue
			}
			alt := dist[item] + neighbor.Weight
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				neighbor.Priority = alt
				heap.Push(queue, neighbor)
			}
		}
	}

	return dist
}
