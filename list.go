package bst

// PreOrder is a public method that calls the private method preOrder on the
// root node of the provided tree
func (tree *Tree) PreOrder() (items []int) {
	items = tree.root.preOrder()
	return
}

// preOrder is a private method that traverses the given node in pre order,
// or the node value followed by the left child followed by the right child
func (node *Node) preOrder() (items []int) {
	if (node == nil) {
		return
	}

	items = append(items, node.value)
	items = append(items, node.left.preOrder()...)
	items = append(items, node.right.preOrder()...)

	return
}

// InOrder is a public method that calls the private method inOrder on the
// root node of the provided tree
func (tree *Tree) InOrder() (items []int) {
	items = tree.root.inOrder()
	return
}

// inOrder is a private method that traverses the given node in order, or the
// left child followed by the node value followed by the right child
func (node *Node) inOrder() (items []int) {
	if (node == nil) {
		return
	}

	items = append(items, node.left.inOrder()...)
	items = append(items, node.value)
	items = append(items, node.right.inOrder()...)

	return
}

// PostOrder is a public method that calls the private method postOrder on
// the root node of the provided tree.
func (tree *Tree) PostOrder() (items []int) {
	items = tree.root.postOrder()
	return
}

// postOrder is a private method that traverses the given node in post order,
// or the left child followed by the right child followed by the node value
func (node *Node) postOrder() (items []int) {
	if (node == nil) {
		return
	}

	items = append(items, node.left.postOrder()...)
	items = append(items, node.right.postOrder()...)
	items = append(items, node.value)

	return
}