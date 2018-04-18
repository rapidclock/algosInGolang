package graphs

import (
	"math"
	"../genutils"
)

// Ford Fulkerson Algorithm to find Maximum Flow in a Weighted Directed Graph via Edmunds Karp Method.
func FordFulkersonMaxFlow(g *WeightedGraph, source, sink Vertex) ([]genutils.Tuple, int) {
	residualNetwork := buildResidualNetwork(g)
	maxFlow := 0
	var augmentedPaths = make([]genutils.Tuple, 0)
	for ok, parent := findAugmentedPath(residualNetwork, source, sink);
		ok; ok, parent = findAugmentedPath(residualNetwork, source, sink) {
		path, minFlowOnPath := augmentPath(residualNetwork, parent, sink)
		//fmt.Println(minFlowOnPath)
		maxFlow += minFlowOnPath
		augmentedPaths = append(augmentedPaths, genutils.Tuple{X: path, Y: minFlowOnPath})
	}
	return augmentedPaths, maxFlow
}

func buildResidualNetwork(g *WeightedGraph) *WeightedGraph {
	residualNetwork := NewWeightedDirectedGraph()
	for edge, weight := range g.Weights {
		residualNetwork.AddWeightedDirectedEdge(edge.Source, edge.Destination, weight)
	}
	for edge := range g.Edges {
		if !residualNetwork.Edges[Edge{edge.Destination, edge.Source}] {
			residualNetwork.AddWeightedDirectedEdge(edge.Destination, edge.Source, 0)
		}
	}
	return residualNetwork
}

func findAugmentedPath(g *WeightedGraph, source, sink Vertex) (bool, map[Vertex]Vertex) {
	if source == nil || sink == nil {
		return false, nil
	}
	parent := make(map[Vertex]Vertex)
	queue := make([]Vertex, 0)
	queue = append(queue, source)
	parent[source] = nil
	for len(queue) > 0 {
		var node Vertex
		node, queue = queue[0], queue[1:]
		for _, neigh := range g.Neighbours[node] {
			edge := Edge{node, neigh}
			if _, ok := parent[neigh]; !ok && g.Weights[edge] > 0 {
				parent[neigh] = node
				queue = append(queue, neigh)
				if neigh == sink {
					return true, parent
				}
			}
		}
	}
	return false, nil
}

func augmentPath(g *WeightedGraph, parent map[Vertex]Vertex, sink Vertex) ([]Vertex, int) {
	var minFlowOnPath = math.MaxInt64
	var node = sink
	var path = make([]Vertex, 0)
	path = append(path, sink)
	for parent[node] != nil {
		parentNode := parent[node]
		weight := g.Weights[Edge{parentNode, node}]
		minFlowOnPath = genutils.MinOfTwo(minFlowOnPath, weight)
		path = append(path, parentNode)
		node = parentNode
	}
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
	for i := 0; i < len(path)-1; i++ {
		edge := Edge{path[i], path[i+1]}
		reverseEdge := Edge{path[i+1], path[i]}
		g.Weights[edge] -= minFlowOnPath
		g.Weights[reverseEdge] += minFlowOnPath
	}
	return path, minFlowOnPath
}
