package main

import (
	"sync"
)

//Node : Base structure for graph node implementation.
type Node struct {
	mCar          *Car
	id            string
	permit        string
	isSemaphor    bool
	semaphorState bool
	hasCar        bool
	isFinal       bool
}

func (n *Node) getCar() *Car {
	return n.mCar
}

func (n *Node) setCar(mCar *Car) {
	n.mCar = mCar
}

func (n *Node) getIsSemaphor() bool {
	return n.isSemaphor
}

func (n *Node) setIsSemaphor(isSemaphor bool) {
	n.isSemaphor = isSemaphor
}

func (n *Node) getSemaphorState() bool {
	return n.semaphorState
}

func (n *Node) setSemaphorState(semaphoreState bool) {
	n.semaphorState = semaphoreState
}

func (n *Node) getHasCar() bool {
	return n.hasCar
}

func (n *Node) setHasCar(hasCar bool) {
	n.hasCar = hasCar
}

func (n *Node) getisFinal() bool {
	return n.isFinal
}

//ItemGraph : Node storage (graph).
type ItemGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

//AddNode : used to add a node to the graph.
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

//AddEdge : addes an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)

	g.lock.Unlock()
}
