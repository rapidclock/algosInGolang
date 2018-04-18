// General Structure for Graphs and Graph Algorithms.
package graphs

import "fmt"

type Vertex interface{}
type Edge struct {
	Source, Destination Vertex
}

type Graph struct {
	Vertices   map[Vertex]bool
	Edges      map[Edge]bool
	Neighbours map[Vertex][]Vertex
}

type WeightedGraph struct {
	Graph
	Weights map[Edge]int
}

func NewGraph() *Graph {
	g := new(Graph)
	g.Vertices = make(map[Vertex]bool)
	g.Edges = make(map[Edge]bool)
	g.Neighbours = make(map[Vertex][]Vertex)
	return g
}

func NewWeightedDirectedGraph() *WeightedGraph {
	wg := new(WeightedGraph)
	wg.Graph = *NewGraph()
	wg.Weights = make(map[Edge]int)
	return wg
}

func (g *Graph) AddEdge(source, destination Vertex) {
	g.AddVertex(source)
	g.AddVertex(destination)
	g.Neighbours[source] = append(g.Neighbours[source], destination)
	edge := Edge{source, destination}
	g.Edges[edge] = true
}

func (g *Graph) AddVertex(v Vertex) {
	g.Vertices[v] = true
}

func (g *Graph) String() string {
	return fmt.Sprintf("%v", g.Neighbours)
}

func (g *WeightedGraph) AddWeightedDirectedEdge(source, destination Vertex, weight int) {
	g.AddEdge(source, destination)
	g.Weights[Edge{Source: source, Destination: destination}] = weight
}

func (g *WeightedGraph) String() string {
	return fmt.Sprintf("Neighbours : %v\nWeights: %v", g.Neighbours, g.Weights)
}
