package bst

// A Node represents the value of a single node and the left and right
// children of that node
type Node struct {
	value int
	left *Node
	right *Node
}

// A Tree represents the root node and the number of nodes in a BST
type Tree struct {
	root *Node
	nodes int
}

// Nodes returns the number of nodes in a tree
func (tree *Tree) Nodes() int {
	return tree.nodes
}

// Root returns the root node of a tree
func (tree *Tree) Root() *Node {
	return tree.root
}

// NewBST initializes a new BST
func NewBST() *Tree {
	tree := new(Tree)
	tree.nodes = 0

	return tree
}