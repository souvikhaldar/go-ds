package graph

import (
	"github.com/souvikhaldar/go-ds/queue"
	"github.com/souvikhaldar/go-ds/stack"
)

type AMGraph struct {
	AdjacencyMatrix [][]string
	vertexMap       map[string]int
	reverseMap      map[int]string
	vertexCount     int
}

// NewAMGraph create a graph which will be stored in
// adjancency matrix of mr rows and mc columns
func NewAMGraph(mr, mc int) *AMGraph {
	am := make([][]string, mr)
	for i := range am {
		am[i] = make([]string, mc)
	}
	return &AMGraph{
		AdjacencyMatrix: am,
		vertexMap:       make(map[string]int),
		reverseMap:      make(map[int]string),
		vertexCount:     0,
	}
}

func (a *AMGraph) AddNode(
	startNodeName string,
	newNodeName string,
	relationShip string,
) {
	var sn, nn int
	sn, ok := a.getNodeInt(startNodeName)
	if !ok {
		sn = a.setNodeName(startNodeName)
	}

	nn, ok = a.getNodeInt(newNodeName)
	if !ok {
		nn = a.setNodeName(newNodeName)
	}

	// TODO: check if it's a valid relationship

	a.AdjacencyMatrix[sn][nn] = relationShip
}

func (a *AMGraph) DeleteNode(nodeName string) {
	node, ok := a.getNodeInt(nodeName)
	if !ok {
		// node not present in graph
		return
	}
	for i := 0; i <= a.vertexCount; i++ {
		a.AdjacencyMatrix[node][i] = ""
		a.AdjacencyMatrix[i][node] = ""

	}
	delete(a.vertexMap, nodeName)
}

func (a *AMGraph) SearchNode(nodeName string) bool {
	_, ok := a.getNodeInt(nodeName)
	return ok
}

func (a *AMGraph) GetNeighboursOf(nodeName string) []string {
	node, ok := a.getNodeInt(nodeName)
	if !ok {
		// node not present in graph
		return nil
	}
	var neighbours []string
	for i := 0; i <= a.vertexCount; i++ {
		if a.AdjacencyMatrix[node][i] != "" {
			nd := a.nodetoName(i)
			neighbours = append(neighbours, nd)

		}
	}
	return neighbours
}

//We explore nodes of the graph layer by layer, meaning we explore all the nodes connected to a node, and
// then proceed further. This becomes useful in finding the shortest path between the nodes. It can also help
// find all the **Connected Componenets** (i.e a subgraph in which all nodes are connected) in a graph.
func (a *AMGraph) BFS(startNode string) []string {
	q := queue.NewQueue()
	explored := make(map[string]bool)
	q.Enqueue(startNode)
	nodes := make([]string, 0)

	for !q.IsEmpty() {
		n, err := q.Dequeue()
		if err != nil {
			panic(err)
		}
		node := n.(string)
		if _, ok := explored[node]; !ok {
			nodes = append(nodes, node)
			// add this node's neighbours to be explored later
			nb := a.GetNeighboursOf(node)
			for _, n := range nb {
				q.Enqueue(n)
			}
			// make this node as explored
			explored[node] = true
		}
	}
	return nodes

}

// DFS (Depth First Search): We explore nodes one after the other without exploring other nodes at the
// same level, and we backtrack only when required, like when no more nodes are left to be explored
// at a given level. For the implementation of this traversal, we use **Stacks** instead of Queue that
// is used in BFS. Time complexity is the same as BFS i.e O(n) where n is the number of nodes, or, linear time complexity.
func (a *AMGraph) DFS(startNode string) []string {
	// keep track of all the explored nodes
	explored := make(map[string]struct{})
	s := stack.NewStack()
	s.Push(startNode)

	var nodes []string
	// keep traversing till stack is not empty
	for !s.IsEmpty() {
		tn, err := s.Pop()
		if err != nil {
			panic(err)
		}
		n, ok := tn.(string)
		if !ok {
			panic("Error in type assertion")
		}
		if _, ok := explored[n]; ok {
			// this node has already been explored
			continue
		}
		// mark it as explored
		explored[n] = struct{}{}
		nodes = append(nodes, n)

		for _, nb := range a.GetNeighboursOf(n) {
			s.Push(nb)
		}
	}
	return nodes
}

func (a *AMGraph) GetAllNodes() []string {
	var nodes []string
	for n, _ := range a.vertexMap {
		nodes = append(nodes, n)
	}
	return nodes
}

// GetRelationBetween returns the relationship between the two given
// nodes.
// It returns empty string if there is no relation
func (a *AMGraph) GetRelationBetween(firstNode, secondNode string) string {
	first, ok := a.getNodeInt(firstNode)
	if !ok {
		return ""
	}
	second, ok := a.getNodeInt(secondNode)
	if !ok {
		return ""
	}
	return a.AdjacencyMatrix[first][second]
}
