package tests

import (
	"testing"
	"../graphs"
	"reflect"
)

const TestSuiteName = "Dijkstra's Shortest Path Test : "

func TestDijkstra_EmptyGraph(t *testing.T) {
	graph := graphs.NewWeightedGraph()
	expected := make(map[graphs.Vertex]int)
	actual := graphs.DijkstraShortestPaths(graph, nil)
	if equal := reflect.DeepEqual(actual, expected); !equal {
		t.Errorf(TestSuiteName + "Empty Graph Failure "+
			"\nExpected : %v \nActual : %v\n", expected, actual)
	}
}

func TestDijkstra_SimpleGraph(t *testing.T) {
	graph := graphs.NewWeightedGraph()
	graph.AddWeightedUndirectedEdge(0, 1, 3)
	graph.AddWeightedUndirectedEdge(0, 2, 2)
	graph.AddWeightedUndirectedEdge(0, 3, 4)
	graph.AddWeightedUndirectedEdge(1, 4, 6)
	graph.AddWeightedUndirectedEdge(1, 6, 4)
	graph.AddWeightedUndirectedEdge(2, 4, 5)
	graph.AddWeightedUndirectedEdge(2, 5, 3)
	graph.AddWeightedUndirectedEdge(4, 5, 1)
	graph.AddWeightedUndirectedEdge(3, 4, 3)
	graph.AddWeightedUndirectedEdge(6, 2, 1)
	expected := map[graphs.Vertex]int{0: 0, 1: 3, 2: 2, 3: 4, 4: 6, 5: 5, 6: 3}
	actual := graphs.DijkstraShortestPaths(graph, 0)
	if equal := reflect.DeepEqual(actual, expected); !equal {
		t.Errorf(TestSuiteName + "Simple Graph Failure "+
			"\nExpected : %v \nActual : %v\n", expected, actual)
	}
}

func TestDijkstra_MediumGraph(t *testing.T) {
	graph := graphs.NewWeightedGraph()
	graph.AddWeightedUndirectedEdge(0, 1, 4)
	graph.AddWeightedUndirectedEdge(0, 7, 8)
	graph.AddWeightedUndirectedEdge(1, 7, 11)
	graph.AddWeightedUndirectedEdge(7, 8, 7)
	graph.AddWeightedUndirectedEdge(7, 6, 1)
	graph.AddWeightedUndirectedEdge(1, 2, 8)
	graph.AddWeightedUndirectedEdge(2, 8, 2)
	graph.AddWeightedUndirectedEdge(8, 6, 6)
	graph.AddWeightedUndirectedEdge(2, 5, 4)
	graph.AddWeightedUndirectedEdge(2, 3, 7)
	graph.AddWeightedUndirectedEdge(6, 5, 2)
	graph.AddWeightedUndirectedEdge(3, 5, 14)
	graph.AddWeightedUndirectedEdge(3, 4, 9)
	graph.AddWeightedUndirectedEdge(5, 4, 10)
	expected := map[graphs.Vertex]int{0: 0, 1: 4, 2: 12, 3: 19, 4: 21, 5: 11, 6: 9, 7: 8, 8: 14}
	actual := graphs.DijkstraShortestPaths(graph, 0)
	if equal := reflect.DeepEqual(actual, expected); !equal {
		t.Errorf(TestSuiteName + "Medium Graph "+
			"\nExpected : %v \nActual : %v\n", expected, actual)
	}
}

func TestDijkstra_18NodeGraph(t *testing.T) {
	graph := graphs.NewWeightedGraph()
	graph.AddWeightedUndirectedEdge(0, 1, 9)
	graph.AddWeightedUndirectedEdge(0, 7, 7)
	graph.AddWeightedUndirectedEdge(0, 14, 6)
	graph.AddWeightedUndirectedEdge(1, 2, 5)
	graph.AddWeightedUndirectedEdge(2, 6, 9)
	graph.AddWeightedUndirectedEdge(3, 6, 5)
	graph.AddWeightedUndirectedEdge(3, 10, 9)
	graph.AddWeightedUndirectedEdge(3, 17, 1)
	graph.AddWeightedUndirectedEdge(4, 5, 9)
	graph.AddWeightedUndirectedEdge(5, 6, 1)
	graph.AddWeightedUndirectedEdge(5, 8, 9)
	graph.AddWeightedUndirectedEdge(5, 9, 2)
	graph.AddWeightedUndirectedEdge(6, 9, 4)
	graph.AddWeightedUndirectedEdge(7, 14, 7)
	graph.AddWeightedUndirectedEdge(7, 11, 7)
	graph.AddWeightedUndirectedEdge(7, 8, 5)
	graph.AddWeightedUndirectedEdge(8, 11, 2)
	graph.AddWeightedUndirectedEdge(8, 9, 7)
	graph.AddWeightedUndirectedEdge(9, 12, 3)
	graph.AddWeightedUndirectedEdge(9, 10, 9)
	graph.AddWeightedUndirectedEdge(10, 13, 8)
	graph.AddWeightedUndirectedEdge(10, 17, 4)
	graph.AddWeightedUndirectedEdge(11, 14, 5)
	graph.AddWeightedUndirectedEdge(11, 15, 8)
	graph.AddWeightedUndirectedEdge(11, 12, 9)
	graph.AddWeightedUndirectedEdge(12, 15, 4)
	graph.AddWeightedUndirectedEdge(12, 16, 7)
	graph.AddWeightedUndirectedEdge(13, 17, 4)
	graph.AddWeightedUndirectedEdge(14, 15, 4)
	graph.AddWeightedUndirectedEdge(15, 17, 2)
	expected := map[graphs.Vertex]int{0: 0, 1: 9, 2: 14, 3: 13, 4: 28, 5: 19, 6: 18, 7: 7, 8: 12, 9: 17,
		10: 16, 11: 11, 12: 14, 13: 16, 14: 6, 15: 10, 16: 21, 17: 12}
	actual := graphs.DijkstraShortestPaths(graph, 0)
	if equal := reflect.DeepEqual(actual, expected); !equal {
		t.Errorf(TestSuiteName + "18 Node Graph Failure "+
			"\nExpected : %v \nActual : %v\n", expected, actual)
	}
}
