package bst

// Delete is a public method that returns true if a value is removed from
// a tree, or false if that value is not in the tree.
func (tree *Tree) Delete(value int) bool {
	// Check if root is nil, which means the tree is empty
	if tree.root == nil {
		return false
	}

	if tree.root.value == value {
		// Tree root is equal to the value being deleted

		// Create new temprorary root so delete can be called and the real
		// root will have a parent if necessary
		tempRoot := &Node{0, nil, nil}

		// Link the real root to the left child of the temprorary root
		tempRoot.left = tree.root

		// Call delete
		result := tree.root.delete(tempRoot, value)

		// Reset the tree root
		tree.root = tempRoot.left

		return result
	}

	// Root is not equal to the value being deleted, try left / right subtrees
	if (!tree.root.left.delete(tree.root, value)) {
		return tree.root.right.delete(tree.root, value)
	}

	// This return statement is reached only if deleting from the left subtree
	// was successful
	return true
}

// delete is a private method that is recursively called to remove a value
// from a given node's subtrees.
func (node *Node) delete(parent *Node, value int) bool {
	// Check that the node exists
	if (node == nil) {
		return false
	}

	if (value > node.value) {
		// Value being deleted is greater than the value of the current node

		if (node.right == nil) {
			// If node has no right subtree, value being deleted is not in
			// the tree
			return false
		}

		// Value being deleted could be in right subtree, call delete on it
		return node.right.delete(node, value)
	} else if (value < node.value) {
		// Value being deleted is less than the value of the current node

		if node.left == nil {
			// If node has no left subtree, value being deleted is not in
			// the tree
			return false
		}

		// Value being deleted could be in left subtree, call delete on it
		return node.left.delete(node, value)
	} else if (node.value == value) {
		// Correct node to delete found

		if (node.left != nil && node.right != nil) {
			// Node has left and right children/subtrees

			// Find minimum value of right subtree; set this node's value
			// to the value found, which will duplicate that value
			node.value = node.right.minValue()

			// Call delete to remove the duplicate node at the bottom of the
			// right subtree
			return node.right.delete(node, node.value)
		}

		// Connect this node's children to its parent
		parent.connectNode(node)

		return true
	}

	return false
}

// minValue is a private method that is recursively called to find the
// minimum value from a given node's subtrees.
func (node *Node) minValue() int {
	if node.left == nil {
		// If this node has no left child, it is the lowest value in the
		// subtree it exists in
		return node.value
	}

	// Call minValue recursively on the left subtree
	return node.left.minValue()
}

// connectNodes is a private method that removes a node from the tree by
// by linking a parent's children to the given node's children.
func (parent *Node) connectNode(node *Node) {
	if parent.left == node {
		if node.left != nil {
			parent.left = node.left
		} else {
			parent.left = node.right
		}
	} else if parent.right == node {
		if node.left != nil {
			parent.right = node.left
		} else {
			parent.right = node.right
		}
	}
}