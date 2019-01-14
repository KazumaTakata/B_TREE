package Graph

import "sort"

type Edge struct {
	StartID ID
	EndID   ID
	Weight  int
}

func MST(g Graph) []Edge {
	unions := []map[ID]bool{}
	unions = append(unions, map[ID]bool{})
	mstSet := []Edge{}
	answer := []Edge{}

	for startNodeId := range g.Weight {
		for endNodeId := range g.Weight[startNodeId] {
			weight := g.Weight[startNodeId][endNodeId]
			mstSet = append(mstSet, Edge{StartID: startNodeId, EndID: endNodeId, Weight: weight})
		}
	}

	sort.Slice(mstSet, func(i, j int) bool {
		return mstSet[i].Weight < mstSet[j].Weight
	})

	for _, edge := range mstSet {
		si := findSet(&unions, edge.StartID)
		ei := findSet(&unions, edge.EndID)
		if si == -1 || ei == -1 || si == ei {
			answer = append(answer, edge)
			margeSet(&unions, edge.StartID, edge.EndID)
		}

	}
	return answer
}

func findSet(unions *[]map[ID]bool, id ID) int {
	for index, unionMap := range *unions {
		_, ok := unionMap[id]
		if ok {
			return index
		}
	}
	return -1
}

func margeSet(unions *[]map[ID]bool, startId ID, endId ID) {
	si := findSet(unions, startId)
	ei := findSet(unions, endId)

	if si == -1 && ei != -1 {
		(*unions)[ei][startId] = true
	}

	if si != -1 && ei == -1 {
		(*unions)[si][endId] = true
	}

	if si != -1 && ei != -1 {
		(*unions)[si][endId] = true
		s_union := map[ID]bool{}
		for k, _ := range (*unions)[si] {
			s_union[k] = true
		}
		for k, _ := range (*unions)[ei] {
			s_union[k] = true
		}

		if si > ei {
			(*unions) = append((*unions)[:si], (*unions)[si+1:]...)
			(*unions) = append((*unions)[:ei], (*unions)[ei+1:]...)
		} else {
			(*unions) = append((*unions)[:ei], (*unions)[ei+1:]...)
			(*unions) = append((*unions)[:si], (*unions)[si+1:]...)
		}
		(*unions) = append((*unions), s_union)
	}

}
