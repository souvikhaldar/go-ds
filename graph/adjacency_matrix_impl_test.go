package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	g := NewAMGraph(8, 8)
	g.AddNode("a", "b", "brother")
	g.AddNode("a", "c", "sister")
	g.AddNode("a", "d", "son")
	g.AddNode("b", "x", "son")
	g.AddNode("c", "y", "son")
	g.AddNode("d", "z", "son")
	g.DeleteNode("b")
	nodes := g.BFS("a")
	if len(nodes) == 0 {
		t.Fatal("BFS failing: ", nodes)
	}
	if g.SearchNode("b") {
		t.Fatal("b should not be present")
	}
	if len(g.GetAllNodes()) == 0 {
		t.Fatal("Failed to get all nodes")
	}
	if g.GetRelationBetween("a", "c") == "" {
		t.Fatal("Relation missing")
	}
}

func TestDFS(t *testing.T) {
	g := NewAMGraph(7, 7)
	g.AddNode("s", "a", "connected")
	g.AddNode("s", "b", "connected")
	g.AddNode("a", "c", "connected")
	g.AddNode("a", "b", "connected")
	g.AddNode("a", "s", "connected")
	g.AddNode("c", "a", "connected")
	g.AddNode("c", "d", "connected")
	g.AddNode("c", "e", "connected")
	g.AddNode("e", "c", "connected")
	g.AddNode("e", "d", "connected")
	g.AddNode("d", "e", "connected")
	g.AddNode("d", "c", "connected")
	g.AddNode("d", "b", "connected")
	g.AddNode("b", "d", "connected")
	g.AddNode("b", "a", "connected")
	g.AddNode("b", "s", "connected")
	nodes := g.DFS("s")
	if len(nodes) == 0 {
		t.Log("Nodes are more than 0 rather.")
	}
}
