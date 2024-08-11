package datastructures

type Graph struct {
	v   []Vertex
	e   []Edge
	adj Adj
}

type Adj []AdjNode

type AdjNode struct {
	Next     *AdjNode
	VertexId int
}

type Vertex struct {
	Id       int
	Color    Color
	Distance int
	Parent   int
}

type Edge struct {
	LeftId, RightId int
}

type Color int

const (
	WHITE Color = iota
	GRAY
	BLACK
)

func CreateGraph(vertexes []int, edges []Edge) *Graph {
	g := &Graph{
		v: createVertexArray(vertexes),
		e: edges,
	}

	return g
}

func createVertexArray(vertexes []int) (v []Vertex) {
	v = make([]Vertex, 0)
	for _, i := range vertexes {
		v[i] = Vertex{
			Id:       i,
			Color:    WHITE,
			Distance: -1,
			Parent:   -1,
		}
	}

	return v
}

// func createAdjacencyList(vertexes []Vertex, edges []Edges) (adj Adj) {

// }

// Add a vertex with its adjacents in an array
func (g Graph) InsertArray(vertex int, adjacents []int) {
	if len(adjacents) == 0 {
		return
	}

	g.adj[vertex] = AdjNode{
		Next:     nil,
		VertexId: adjacents[0],
	}

	for i := 1; i < len(adjacents); i++ {
		g.adj[vertex].Next = &AdjNode{
			Next:     nil,
			VertexId: adjacents[i],
		}
	}
}
