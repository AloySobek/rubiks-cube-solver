package graph

import (
	"bytes"
	"math/rand"
	"os"

	"github.com/AloySobek/Rubik/cube"
	"github.com/bradleyjkemp/memviz"
)

type Label int64

const (
	ORIGIN       = iota
	INTERMEDIATE = iota
	GOAL         = iota
)

type Node struct {
	Data  *Data
	Edges []*Edge
}

type Data struct {
	Cube     *cube.Cube
	Path     *Node
	Distance int64
	Label    Label
}

type Edge struct {
	Meta *Meta
	Node *Node
}

type Meta struct {
	Weight int64
	Move   rune
}

func CreateNode(data *Data) *Node {
	return &Node{
		Data:  data,
		Edges: make([]*Edge, 0),
	}
}

func CreateEdge(dst *Node, meta *Meta) *Edge {
	return &Edge{
		Node: dst,
		Meta: meta,
	}
}

func AttachEdge(node *Node, edge *Edge) *Node {
	node.Edges = append(node.Edges, edge)

	return node
}

func GetRandomGraph(origin, goal *Node) *Node {
	if origin == nil {
		origin = CreateNode(&Data{
			Cube:     nil,
			Path:     nil,
			Distance: 0,
			Label:    ORIGIN,
		})
	}

	if goal == nil {
		goal = CreateNode(&Data{
			Cube:     nil,
			Path:     nil,
			Distance: int64(^uint(0) >> 1),
			Label:    GOAL,
		})
	}

	GenerateRandomIntermidateNodes(origin, goal, rand.Intn(2)+3)

	return origin
}

func GenerateRandomIntermidateNodes(origin, goal *Node, depth int) *Node {
	if origin == nil || goal == nil {
		panic("Origin and Goal nodes must be presented!")
	}

	if depth <= 0 {
		return origin
	} else if depth == 1 {
		AttachEdge(origin, &Edge{
			Node: goal,
			Meta: &Meta{
				Weight: int64(rand.Intn(99) + 1),
				Move:   'R',
			},
		})
	} else {
		limit := rand.Intn(2) + 3

		for i := 0; i < limit; i++ {
			intermediate := CreateNode(&Data{
				Cube:     nil,
				Path:     nil,
				Distance: int64(^uint(0) >> 1),
				Label:    INTERMEDIATE,
			})

			AttachEdge(origin, &Edge{
				Node: intermediate,
				Meta: &Meta{
					Weight: int64(rand.Intn(99) + 1),
					Move:   'R',
				},
			})

			GenerateRandomIntermidateNodes(intermediate, goal, depth-1)
		}
	}

	return origin
}

func Print(origin *Node) {
	buf := &bytes.Buffer{}
	memviz.Map(buf, origin)
	err := os.WriteFile("example-tree-data.dot", buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	// switch origin.Data.Label {
	// case ORIGIN:
	// 	fmt.Println("O")
	// case INTERMEDIATE:
	// 	fmt.Println("I")
	// case GOAL:
	// 	fmt.Println("G")
	// }

	// for _, i := range origin.Edges {
	// 	Print(i.Node)
	// }
}

func PremadeGraphOne() *Node {
	one := CreateNode(&Data{
		Cube:     nil,
		Path:     nil,
		Distance: 0,
		Label:    ORIGIN,
	})

	two := CreateNode(&Data{
		Cube:     nil,
		Path:     nil,
		Distance: int64(^uint(0) >> 1),
		Label:    INTERMEDIATE,
	})

	three := CreateNode(&Data{
		Cube:     nil,
		Path:     nil,
		Distance: int64(^uint(0) >> 1),
		Label:    INTERMEDIATE,
	})

	four := CreateNode(&Data{
		Cube:     nil,
		Path:     nil,
		Distance: int64(^uint(0) >> 1),
		Label:    INTERMEDIATE,
	})

	five := CreateNode(&Data{
		Cube:     nil,
		Path:     nil,
		Distance: int64(^uint(0) >> 1),
		Label:    GOAL,
	})

	six := CreateNode(&Data{
		Cube:     nil,
		Path:     nil,
		Distance: int64(^uint(0) >> 1),
		Label:    INTERMEDIATE,
	})

	AttachEdge(one, &Edge{
		Node: two,
		Meta: &Meta{
			Weight: 7,
			Move:   'R',
		},
	})

	AttachEdge(one, &Edge{
		Node: three,
		Meta: &Meta{
			Weight: 9,
			Move:   'R',
		},
	})

	AttachEdge(one, &Edge{
		Node: six,
		Meta: &Meta{
			Weight: 14,
			Move:   'R',
		},
	})

	AttachEdge(two, &Edge{
		Node: three,
		Meta: &Meta{
			Weight: 10,
			Move:   'R',
		},
	})

	AttachEdge(two, &Edge{
		Node: four,
		Meta: &Meta{
			Weight: 15,
			Move:   'R',
		},
	})

	AttachEdge(three, &Edge{
		Node: four,
		Meta: &Meta{
			Weight: 11,
			Move:   'R',
		},
	})

	AttachEdge(three, &Edge{
		Node: six,
		Meta: &Meta{
			Weight: 2,
			Move:   'R',
		},
	})

	AttachEdge(four, &Edge{
		Node: five,
		Meta: &Meta{
			Weight: 6,
			Move:   'R',
		},
	})

	AttachEdge(six, &Edge{
		Node: five,
		Meta: &Meta{
			Weight: 9,
			Move:   'R',
		},
	})

	return one
}
