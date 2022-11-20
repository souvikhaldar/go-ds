package graph

type Graph interface {
	AddNode(startNodeName string, newNodeName string, relationShip string)
	DeleteNode(nodeName string)
	SearchNode(nodeName string) bool
	GetNeighboursOf(nodeName string) []string
	BFS(startNode string) []string
}
