package graphs

import (
	"math"
)

func DijkstraShortestPaths(graph *WeightedGraph, startVertex Vertex) map[Vertex]int {
	distance := make(map[Vertex]int)
	zSet := make(map[Vertex]bool)
	for v := range graph.Vertices {
		if v == startVertex {
			distance[v] = 0
		} else {
			distance[v] = math.MaxInt64
		}
	}
	for len(zSet) != len(graph.Vertices) {
		vertex := getMinDistNotInSet(distance, zSet)
		zSet[vertex] = true
		for _, neighbour := range graph.Neighbours[vertex] {
			if !zSet[neighbour] {
				relaxEdge(vertex, neighbour, graph, distance)
			}
		}
	}
	return distance
}

func getMinDistNotInSet(distance map[Vertex]int, set map[Vertex]bool) Vertex {
	var minDistVertex Vertex
	var minDist = math.MaxInt64
	for v, dist := range distance {
		if !set[v] && dist < minDist {
			minDistVertex, minDist = v, dist
		}
	}
	return minDistVertex
}

func relaxEdge(u, v Vertex, graph *WeightedGraph, distance map[Vertex]int) {
	edge := Edge{u, v}
	if distance[u]+graph.Weights[edge] < distance[v] {
		distance[v] = distance[u] + graph.Weights[edge]
	}
}
