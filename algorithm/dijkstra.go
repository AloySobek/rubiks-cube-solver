package algorithm

import (
	"container/heap"

	"github.com/AloySobek/Rubik/graph"
)

type NodeHeap []*graph.Node

func (h *NodeHeap) Len() int           { return len(*h) }
func (h *NodeHeap) Less(i, j int) bool { return (*h)[i].Data.Distance < (*h)[j].Data.Distance }
func (h *NodeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*graph.Node))
}
func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Dijkstra(origin *graph.Node) *graph.Node {
	visited, h := map[*graph.Node]bool{}, &NodeHeap{}

	heap.Init(h)

	heap.Push(h, origin)

	for true {
		if len(*h) == 0 {
			break
		}

		node := heap.Pop(h).(*graph.Node)

		if node.Data.Label == graph.GOAL {
			return node
		}

		for _, i := range node.Edges {
			if _, ok := visited[node]; ok {
				continue
			}

			if i.Node.Data.Distance > node.Data.Distance+i.Meta.Weight {
				i.Node.Data.Distance = node.Data.Distance + i.Meta.Weight
				i.Node.Data.Path = node
			}

			heap.Push(h, i.Node)
		}

		visited[node] = true
	}

	return nil
}
