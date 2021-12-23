package dijkstra

import (
	"container/heap"
	"math"
)

// Node is a graph node used in Dijkstra's algorithm.
type Node struct {
	Weight   int
	Distance int
}

// Dijkstra performs Dijkstra's algorithm on the given graph, setting each
// node's "Distance" property to the length of the shortest path from
// the source node.
func Dijkstra(graph map[*Node][]*Node, source *Node) {
	// Initialize distances to "inifinity"
	for node := range graph {
		node.Distance = math.MaxInt64
	}
	source.Distance = 0

	// Initialize queue
	queue := &minPQ{}
	heap.Init(queue)
	heap.Push(queue, source)

	visited := make(map[*Node]bool)

	// Compute
	for queue.Len() > 0 {
		item := heap.Pop(queue).(*Node)
		visited[item] = true

		for _, neighbor := range graph[item] {
			if visited[neighbor] {
				continue
			}
			alt := item.Distance + neighbor.Weight
			if alt < neighbor.Distance {
				neighbor.Distance = alt
				heap.Push(queue, neighbor)
			}
		}
	}
}

// Minimum Priority Queue
type minPQ []*Node

func (pq minPQ) Len() int { return len(pq) }

func (pq minPQ) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
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
