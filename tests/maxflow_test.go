package tests

import (
	"testing"
	"../graphs"
)

func TestMaxFlow_EmptyGraph(t *testing.T) {
	g := graphs.NewWeightedGraph()
	paths, actualMaxFlow := graphs.FordFulkersonMaxFlow(g, nil, nil)
	expectedMaxFlow, expectedNumOfPaths := 0, 0
	if actualMaxFlow != expectedMaxFlow {
		t.Errorf("Max Flow: Empty Graph Case Failed, Max Flow Incorrect. "+
			"Expected : %v, Actual : %v", expectedMaxFlow, actualMaxFlow)
	}
	if expectedNumOfPaths != len(paths) {
		t.Errorf("Max Flow: Empty Graph Case Failed, "+
			"Number of Augmented Paths Incorrect. Expected : %v, Actual : %v",
			expectedNumOfPaths, len(paths))
	}
}

func TestMaxFlow_SimpleAdversarialCase(t *testing.T) {
	g := graphs.NewWeightedGraph()
	g.AddWeightedDirectedEdge(0, 1, 1000)
	g.AddWeightedDirectedEdge(0, 2, 1000)
	g.AddWeightedDirectedEdge(2, 1, 1)
	g.AddWeightedDirectedEdge(2, 3, 1000)
	g.AddWeightedDirectedEdge(1, 3, 1000)
	paths, actualMaxFlow := graphs.FordFulkersonMaxFlow(g, 0, 3)
	expectedMaxFlow, expectedNumOfPaths := 2000, 2
	if actualMaxFlow != expectedMaxFlow {
		t.Errorf("Max Flow: Simple Adversarial Case Failed, Max Flow Incorrect. "+
			"Expected : %v, Actual : %v", expectedMaxFlow, actualMaxFlow)
	}
	if expectedNumOfPaths != len(paths) {
		t.Errorf("Max Flow: Simple Adversarial Case Failed, "+
			"Number of Augmented Paths Incorrect. Expected : %v, Actual : %v",
			expectedNumOfPaths, len(paths))
	}
}

func TestMaxFlow_MediumExampleStringVertex(t *testing.T) {
	g := graphs.NewWeightedGraph()
	g.AddWeightedDirectedEdge("S", "b", 4)
	g.AddWeightedDirectedEdge("b", "a", 2)
	g.AddWeightedDirectedEdge("b", "c", 1)
	g.AddWeightedDirectedEdge("b", "d", 2)
	g.AddWeightedDirectedEdge("a", "S", 4)
	g.AddWeightedDirectedEdge("a", "d", 3)
	g.AddWeightedDirectedEdge("c", "d", 2)
	g.AddWeightedDirectedEdge("c", "T", 3)
	g.AddWeightedDirectedEdge("d", "T", 4)
	paths, actualMaxFlow := graphs.FordFulkersonMaxFlow(g, "S", "T")
	expectedMaxFlow, expectedNumOfPaths := 4, 3
	if actualMaxFlow != expectedMaxFlow {
		t.Errorf("Max Flow: Medium Case with String Vertices Failed, "+
			"Max Flow Incorrect. Expected : %v, Actual : %v", expectedMaxFlow, actualMaxFlow)
	}
	if expectedNumOfPaths != len(paths) {
		t.Errorf("Max Flow: Medium Case with String Vertices Failed, "+
			"Number of Augmented Paths Incorrect. Expected : %v, Actual : %v",
			expectedNumOfPaths, len(paths))
	}
}

func TestMaxFlow_MediumExampleWithBackEdge(t *testing.T) {
	g := graphs.NewWeightedGraph()
	g.AddWeightedDirectedEdge(0, 1, 16)
	g.AddWeightedDirectedEdge(0, 2, 13)
	g.AddWeightedDirectedEdge(1, 2, 10)
	g.AddWeightedDirectedEdge(2, 1, 4)
	g.AddWeightedDirectedEdge(1, 3, 12)
	g.AddWeightedDirectedEdge(3, 2, 9)
	g.AddWeightedDirectedEdge(2, 4, 14)
	g.AddWeightedDirectedEdge(4, 3, 7)
	g.AddWeightedDirectedEdge(3, 5, 20)
	g.AddWeightedDirectedEdge(4, 5, 4)
	paths, actualMaxFlow := graphs.FordFulkersonMaxFlow(g, 0, 5)
	expectedMaxFlow, expectedNumOfPaths := 23, 3
	if actualMaxFlow != expectedMaxFlow {
		t.Errorf("Max Flow: Medium Case with Back Edge Failed, "+
			"Max Flow Incorrect. Expected : %v, Actual : %v", expectedMaxFlow, actualMaxFlow)
	}
	if expectedNumOfPaths != len(paths) {
		t.Errorf("Max Flow: Medium Case with Back Edge Failed, "+
			"Number of Augmented Paths Incorrect. Expected : %v, Actual : %v",
			expectedNumOfPaths, len(paths))
	}
}
