package Graph

import (
	"math"
)

func BellManFord(g *Graph, initID ID) bool {
	InitializeSingleSource(g, initID)
	for i := 1; i < len(g.Nodes); i++ {
		for startID := range g.Weight {
			for endID := range g.Weight[startID] {
				relax(g, startID, endID)
			}
		}
	}
	for startID := range g.Weight {
		for endID := range g.Weight[startID] {
			if g.Nodes[endID].Distance > g.Nodes[startID].Distance+float64(g.Weight[startID][endID]) {
				return false
			}
		}
	}
	return true
}

func InitializeSingleSource(g *Graph, initID ID) {
	for nodeid := range g.Nodes {
		g.Nodes[nodeid].Distance = math.Inf(0)
		g.Nodes[nodeid].Parent = NoParent
	}
	g.Nodes[initID].Distance = 0
}

func relax(g *Graph, startID ID, endID ID) {
	if g.Nodes[endID].Distance > g.Nodes[startID].Distance+float64(g.Weight[startID][endID]) {
		g.Nodes[endID].Distance = g.Nodes[startID].Distance + float64(g.Weight[startID][endID])
		g.Nodes[endID].Parent = startID
	}
}
