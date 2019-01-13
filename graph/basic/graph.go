package Graph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type ID int

const (
	White = iota
	Gray
	Black
)

var NoParent ID = 0

type Node struct {
	Value      int
	Id         ID
	Color      int
	Parent     ID
	Distance   float64
	FirstTime  int
	SecondTime int
}

type Graph struct {
	Nodes  map[ID]*Node
	Weight map[ID]map[ID]int
}

func (g *Graph) AddNode(node *Node) {
	if g.Nodes == nil {
		g.Nodes = make(map[ID]*Node)
	}
	id := node.Id
	g.Nodes[id] = node
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
	Type   string `json:"type"`
	Size   string `json:"size"`
}

func (g *Graph) AddEdge(startID, endID ID, weight int) {

	if g.Weight == nil {
		g.Weight = make(map[ID]map[ID]int)
	}

	if g.Nodes[startID] == nil {
		log.Fatalf("You tried to add edge between %d and %d.　But, there is no node of id %d", startID, endID, startID)
		// newNode := Node{id: startID}
		// g.AddNode(&newNode)
	}

	if g.Nodes[endID] == nil {
		log.Fatalf("You tried to add edge between %d and %d.　But, there is no node of id %d", startID, endID, endID)
		// newNode := Node{id: endID}
		// g.AddNode(&newNode)
	}

	if g.Weight[startID] == nil {
		g.Weight[startID] = make(map[ID]int)
	}

	g.Weight[startID][endID] = weight

}

func (g *Graph) ShowNode() {
	fmt.Printf("list of nodes in graph\n")
	for key := range g.Nodes {
		fmt.Printf("ID:%d, ", key)
	}
}

func (g *Graph) ShowEdge() {
	fmt.Printf("show of edges in graph\n")
	for startID := range g.Weight {
		fmt.Printf("%d ->", startID)
		endsMap := g.Weight[startID]
		for endID := range endsMap {
			fmt.Printf("%d, ", endID)
		}
		fmt.Printf("\n")
	}
}

func (g *Graph) JsonExport() {
	fmt.Printf("export json file")
	exportGraph := GraphExport{}
	for key := range g.Nodes {
		sid := strconv.Itoa(int(key))
		enode := ExportNode{Id: sid, Label: sid}
		exportGraph.Nodes = append(exportGraph.Nodes, enode)
	}

	for startID := range g.Weight {
		endsMap := g.Weight[startID]
		for endID := range endsMap {
			sid := strconv.Itoa(int(startID))
			eid := strconv.Itoa(int(endID))
			jedge := ExportEdge{Id: sid + "-" + eid, Source: sid, Target: eid, Type: "curvedArrow", Size: "3"}
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
