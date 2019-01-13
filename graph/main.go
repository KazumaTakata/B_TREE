package main

import (
	"fmt"

	"github.com/KazumaTakata/datastructures_algorithms/graph/basic"
)

func main() {
	g := Graph.Graph{}
	node1 := Graph.Node{Id: 1}
	g.AddNode(&node1)
	node2 := Graph.Node{Id: 2}
	g.AddNode(&node2)
	node3 := Graph.Node{Id: 3}
	g.AddNode(&node3)
	node4 := Graph.Node{Id: 4}
	g.AddNode(&node4)
	node5 := Graph.Node{Id: 5}
	g.AddNode(&node5)
	node6 := Graph.Node{Id: 6}
	g.AddNode(&node6)

	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 10)
	g.AddEdge(2, 3, 10)
	g.AddEdge(1, 4, 10)
	g.AddEdge(4, 5, 10)
	g.AddEdge(3, 6, 10)

	g.ShowEdge()
	g.JsonExport()

	Graph.BreadthFirstSearch(g, 1)

	fmt.Printf("end")

}
