package graph

type Data struct {
	Name string
	Age  int
}

type Meta struct {
	Length int
}

type Edge struct {
	Meta *Meta
	Node *Node
}

type Node struct {
	Data  *Data
	Edges []*Edge
}

func Connect(root *Node, child *Node, meta *Meta) *Node {
	root.Edges = append(root.Edges, &Edge{
		Meta: meta,
		Node: child,
	})

	return root
}
