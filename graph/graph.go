package graph

import (
	"github.com/AloySobek/Rubik/cube"
)

type Data struct {
	Cube *cube.Cube
}

type Meta struct {
	HeuristicValue int64
}

type Edge struct {
	Meta *Meta
	Node *Node
}

type Node struct {
	Data  *Data
	Edges []*Edge
}

func CreateNode(data *Data) *Node {
	return &Node{
		Data:  data,
		Edges: make([]*Edge, 0),
	}
}

func Connect(a *Node, b *Node, meta *Meta) *Node {
	a.Edges = append(a.Edges, &Edge{
		Meta: meta,
		Node: b,
	})

	return a
}
