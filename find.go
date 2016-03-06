package bst

// Find is a public method that searches for a value within the specified
// tree.  It returns the results of calling the private find method on the
// root node.
func (tree *Tree) Find(value int) (bool, int) {
	// Call private find method on root node
	return tree.root.find(value)
}

// find is a private method that receives a value to search for starting at
// a specified node, and returns true or false depending on if it is found,
// and the number of levels traversed to reach that conclusion.  This method
// is called recursively.
func (current *Node) find(value int) (found bool, depth int) {

	if (current == nil) {
		return false, 1
	}

	if (current.value == value) {
		// Node found
		return true, 1
	} else if (current.value < value) {
		// Value is greater than current node, move right
		found, depth = current.right.find(value)
		depth++
		return
	} else if (current.value > value) {
		// Value is less than current node, move left
		found, depth = current.left.find(value)
		depth++
		return
	}

	return false, 0
}