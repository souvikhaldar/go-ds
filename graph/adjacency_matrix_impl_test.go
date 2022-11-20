package graph

import "testing"

func TestBFS(t *testing.T) {
	g := NewAMGraph()
	g.AddNode("a", "b", "brother")
	g.AddNode("a", "c", "sister")
	g.AddNode("a", "d", "hijda")
	g.AddNode("b", "x", "hijda")
	g.AddNode("c", "y", "hijda")
	g.AddNode("d", "z", "hijda")
	g.DeleteNode("b")
	nodes := g.BFS("a")
	if len(nodes) == 0 {
		t.Fatal("BFS failing: ", nodes)
	}
	if g.SearchNode("b") {
		t.Fatal("b should not be present")
	}
}
