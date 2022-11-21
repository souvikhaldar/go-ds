package graph

import (
	"github.com/souvikhaldar/go-ds/queue"
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
	// add relationship both ways since it's an
	// undirected graph
	a.AdjacencyMatrix[nn][sn] = relationShip
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
