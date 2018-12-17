package main

type Tree struct {
	Root *Node
}

type Node struct {
	Leaf     bool
	N        int
	key      []int
	deg      int
	children []*Node
}

var Degree = 3

func main() {

	tree := createTree(Degree)

}

func createTree(deg int) *Tree {
	tree := &Tree{}
	tree.Root.Leaf = true
	tree.Root.N = 0
	tree.Root.key = make([]int, deg)
	return tree
}

func insertKeyToTree(tree *Tree, key int) {
	root := tree.Root

	if root.N == 2*Degree-1 {
		newnode := &Node{deg: Degree}
		tree.Root = newnode
		newnode.Leaf = false
		newnode.N = 0
		newnode.children[1] = root
		splitChild(newnode, 1)
		insertKeyToNonFull(newnode, key)
	} else {
		insertKeyToNonFull(root, key)
	}
}

func insertKeyToNonFull(node *Node, key int) {
	if node.Leaf {
		i := 0
		for node.N > i {
			for node.key[i] > key {
				i++
			}
		}
		node.key[i] = key
	} else {
		i := 0
		for node.N > i {
			for node.key[i] > key {
				i++
			}
		}
		insertKeyToNonFull(node.children[i], key)
	}
}

func splitChild(node *Node, index int) {
	newchild := &Node{}
	child := node.children[index]
	newchild.Leaf = child.Leaf
	newchild.N = Degree - 1

	for i := 0; i < Degree-1; i++ {
		newchild.key[i] = child.key[Degree+i]
	}

	if !newchild.Leaf {
		for i := 0; i < Degree; i++ {
			newchild.children[i] = child.children[Degree+i]
		}
	}

	newchild.N = Degree - 1

	for i := node.N + 1; i > index; i-- {
		node.children[i+1] = node.children[i]
	}

	for i := node.N; i > index; i-- {
		node.key[i+1] = node.key[i]
	}

	node.key[index] = child.key[Degree]
	node.N = node.N + 1
}
