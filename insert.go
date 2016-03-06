package bst

// Insert is a public method that attempts to insert a value into
// the specified tree.  It returns true if inserted as the root node,
// or the results of calling insert.
func (tree *Tree) Insert(value int) bool {
	// Create new node which will be put into the tree later
	node := &Node{value, nil, nil}

	// If tree root is nil, assign value to root
	if (tree.root == nil) {
		tree.root = node
		tree.nodes++

		return true
	} else {
		// Make recursive call to private insert method
		if ok := tree.root.insert(node); !ok {
			return false
		}

		tree.nodes++
		return true
	}
}

// insert is a private method that is recursively called to insert
// a node as a child of the current node.
func (current *Node) insert(node *Node) bool {
	if (current.value < node.value) {
		// Check if right child is nil
		if (current.right == nil) {
			// Insert new node as right child
			current.right = node
		} else {
			// Call insert on right child node
			return current.right.insert(node)
		}
	} else if (current.value > node.value) {
		// Check if left child is nil
		if (current.left == nil) {
			// Insert new node as left child
			current.left = node
		} else {
			// Call insert on left child node
			return current.left.insert(node)
		}
	} else {
		// Node already exists, cannot be inserted, return false
		return false
	}

	// Node was successfully inserted, return true
	return true
}