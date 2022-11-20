package graph

import "fmt"

func (a *AMGraph) getNodeInt(name string) (int, bool) {
	n, ok := a.vertexMap[name]
	return n, ok
}

func (a *AMGraph) setNodeName(name string) int {
	a.vertexMap[name] = a.vertexCount
	a.reverseMap[a.vertexCount] = name
	savedPosition := a.vertexCount
	a.vertexCount++
	return savedPosition
}

func (a *AMGraph) nodetoName(i int) string {
	name, ok := a.reverseMap[i]
	if !ok {
		panic(fmt.Sprintln("Not mapped in reverse", i))
		return ""
	}
	return name
}
