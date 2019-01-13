package Graph

import "math"

func BreadthFirstSearch(g *Graph, startID ID) {

	for id := range g.Nodes {
		g.Nodes[id].Color = White
		g.Nodes[id].Distance = math.Inf(0)
	}

	startNode := g.Nodes[startID]
	startNode.Color = Gray
	startNode.Distance = 0
	startNode.Parent = NoParent

	queue := Queue{}
	queue.Push(startID)

	for !queue.IsEmpty() {
		nodeId := queue.Pop()
		node := g.Nodes[nodeId]
		nextNodes := g.Weight[nodeId]
		for nextNodeId := range nextNodes {
			nextNode := g.Nodes[nextNodeId]
			if nextNode.Color == White {
				nextNode.Color = Gray
				nextNode.Distance = node.Distance + 1
				nextNode.Parent = nodeId
				queue.Push(nextNodeId)
			}
		}
		node.Color = Black
	}

}

var time int = 0

func DepthFirstSearch(g *Graph, startID ID) {

	for id := range g.Nodes {
		g.Nodes[id].Color = White
		g.Nodes[id].Parent = NoParent
	}

	for nodeId := range g.Nodes {
		node := g.Nodes[nodeId]
		if node.Color == White {
			DepthFirstSearchVisit(g, node)
		}
	}
}

func DepthFirstSearchVisit(g *Graph, node *Node) {
	time = time + 1
	node.FirstTime = time
	node.Color = Gray

	for nextNodeId := range g.Weight[node.Id] {
		nextNode := g.Nodes[nextNodeId]
		nextNode.Parent = node.Id
		DepthFirstSearchVisit(g, nextNode)
	}
	node.Color = Black
	time = time + 1
	node.SecondTime = time
}
