package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type ID int

type Node struct {
	value int
	id    ID
}

type Graph struct {
	nodes  map[ID]*Node
	weight map[ID]map[ID]int
}

func (g *Graph) AddNode(node *Node) {
	if g.nodes == nil {
		g.nodes = make(map[ID]*Node)
	}
	id := node.id
	g.nodes[id] = node
}

type GraphExport struct {
	Nodes []ExportNode `json:"nodes"`
	Edges []ExportEdge `json:"edges"`
}

type ExportNode struct {
	Id    string `json:"id"`
	Label string `json:"label"`
}

type ExportEdge struct {
	Id     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
}

func (g *Graph) AddEdge(startID, endID ID, weight int) {

	if g.weight == nil {
		g.weight = make(map[ID]map[ID]int)
	}

	if g.nodes[startID] == nil {
		log.Fatalf("You tried to add edge between %d and %d.　But, there is no node of id %d", startID, endID, startID)
		// newNode := Node{id: startID}
		// g.AddNode(&newNode)
	}

	if g.nodes[endID] == nil {
		log.Fatalf("You tried to add edge between %d and %d.　But, there is no node of id %d", startID, endID, endID)
		// newNode := Node{id: endID}
		// g.AddNode(&newNode)
	}

	if g.weight[startID] == nil {
		g.weight[startID] = make(map[ID]int)
	}

	g.weight[startID][endID] = weight

}

func (g *Graph) showNode() {
	fmt.Printf("list of nodes in graph\n")
	for key := range g.nodes {
		fmt.Printf("ID:%d, ", key)
	}
}

func (g *Graph) showEdge() {
	fmt.Printf("show of edges in graph\n")
	for startID := range g.weight {
		fmt.Printf("%d ->", startID)
		endsMap := g.weight[startID]
		for endID := range endsMap {
			fmt.Printf("%d, ", endID)
		}
		fmt.Printf("\n")
	}
}

func (g *Graph) jsonExport() {
	fmt.Printf("export json file")
	exportGraph := GraphExport{}
	for key := range g.nodes {
		sid := strconv.Itoa(int(key))
		enode := ExportNode{Id: sid, Label: sid}
		exportGraph.Nodes = append(exportGraph.Nodes, enode)
	}

	for startID := range g.weight {
		endsMap := g.weight[startID]
		for endID := range endsMap {
			sid := strconv.Itoa(int(startID))
			eid := strconv.Itoa(int(endID))
			jedge := ExportEdge{Id: sid + "-" + eid, Source: sid, Target: eid}
			exportGraph.Edges = append(exportGraph.Edges, jedge)
		}

	}

	b, err := json.Marshal(exportGraph)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	_ = ioutil.WriteFile("output.json", b, 0644)

}

func main() {
	g := Graph{}
	node1 := Node{id: 1}
	g.AddNode(&node1)
	node2 := Node{id: 2}
	g.AddNode(&node2)
	node3 := Node{id: 3}
	g.AddNode(&node3)

	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 10)
	g.AddEdge(2, 3, 10)

	g.showEdge()
	g.jsonExport()

}
