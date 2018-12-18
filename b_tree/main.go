package main

import "fmt"

type Tree struct {
	Root *Node
}

type Node struct {
	Leaf     bool
	N        int
	key      []int
	deg      int
	children []*Node
	Last     bool
}

var Degree = 3

func main() {

	tree := &Tree{}
	b_treeCreate(tree)

	insertKeyToTree(tree, 3)
	insertKeyToTree(tree, 5)
	insertKeyToTree(tree, 1)
	insertKeyToTree(tree, 4)
	insertKeyToTree(tree, 9)
	insertKeyToTree(tree, 10)
	insertKeyToTree(tree, 16)
	insertKeyToTree(tree, 11)
	insertKeyToTree(tree, 15)
	insertKeyToTree(tree, 2)
	insertKeyToTree(tree, 6)
	insertKeyToTree(tree, 7)
	insertKeyToTree(tree, 14)
	insertKeyToTree(tree, 13)
	insertKeyToTree(tree, 17)
	insertKeyToTree(tree, 18)
	insertKeyToTree(tree, 34)
	insertKeyToTree(tree, 23)
	insertKeyToTree(tree, 27)
	insertKeyToTree(tree, 21)
	insertKeyToTree(tree, 22)
	insertKeyToTree(tree, 25)
	insertKeyToTree(tree, 24)
	PrintTree(tree)
}

func allocNode() *Node {
	newnode := &Node{deg: Degree}
	newnode.key = []int{}
	newnode.children = []*Node{}
	return newnode
}

func b_treeCreate(tree *Tree) {
	newnode := allocNode()
	newnode.Leaf = true
	newnode.N = 0
	tree.Root = newnode
}

func insertKeyToTree(tree *Tree, key int) {
	root := tree.Root

	if root.N == 2*Degree-1 {
		newnode := allocNode()
		tree.Root = newnode
		newnode.Leaf = false
		newnode.N = 0
		newnode.children = append(newnode.children[:0], append([]*Node{root}, newnode.children[0:]...)...)
		splitChild(newnode, 0)
		insertKeyToNonFull(newnode, key)
	} else {
		insertKeyToNonFull(root, key)
	}
}

func insertKeyToNonFull(node *Node, key int) {
	if node.Leaf {
		i := 0
		for node.N > i && node.key[i] < key {
			i++
		}
		node.key = append(node.key[:i], append([]int{key}, node.key[i:]...)...)
		node.N = node.N + 1
	} else {
		i := 0
		for node.N > i && node.key[i] < key {
			i++
		}
		if node.children[i].N == 2*Degree-1 {
			splitChild(node, i)
			if key > node.key[i] {
				i = i + 1
			}
		}
		insertKeyToNonFull(node.children[i], key)

	}
}

func splitChild(node *Node, index int) {
	newchild := allocNode()
	child := node.children[index]
	newchild.Leaf = child.Leaf
	newchild.N = Degree - 1

	for i := 0; i < Degree-1; i++ {
		newchild.key = append(newchild.key, child.key[Degree+i])
	}
	if !newchild.Leaf {
		for i := 0; i < Degree; i++ {
			newchild.children = append(newchild.children, child.children[Degree+i])
		}
		child.children = child.children[:Degree]
	}
	child.N = Degree - 1
	centerKey := child.key[child.N]
	child.key = child.key[:child.N]

	for i := node.N; i > index; i-- {
		node.children[i+1] = node.children[i]
	}
	node.children = append(node.children[:index+1], append([]*Node{newchild}, node.children[index+1:]...)...)

	for i := node.N; i > index; i-- {
		node.key[i+1] = node.key[i]
	}

	node.key = append(node.key[:index], append([]int{centerKey}, node.key[index:]...)...)
	node.N = node.N + 1
}

func PrintTree(tree *Tree) {

	treeString := ""

	queue := []*Node{}
	root := tree.Root
	root.Last = true

	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, key := range node.key {
			treeString += fmt.Sprintf("%d ", key)
		}
		if node.Last {
			treeString += "\n"
		} else {
			treeString += "|"
		}

		for i, nc := range node.children {
			if node.Last {
				if i == len(node.children)-1 {
					nc.Last = true
				} else {
					nc.Last = false
				}
			} else {
				nc.Last = false
			}
			queue = append(queue, nc)
		}

	}

	fmt.Println(treeString)

}
