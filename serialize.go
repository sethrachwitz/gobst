package bst

import (
	"os"
	"encoding/json"
)

// Serialize is a public method that serializes a tree to a file, so
// it can be deserialized at a later time into a new tree, and returns
// true if successful and false if not.
func (tree *Tree) Serialize(filename string) bool {
	// Check if filename is an empty string
	if (filename == "") {
		return false
	}

	// Get items from tree in pre order so they can be reinserted
	// in the correct order when deserialized
	items := tree.PreOrder()

	// Create the output file
	file, err := os.Create(filename)
	if (err != nil) {
		return false
	}
	defer file.Close()

	// Create new encoder for the output file
	encoder := json.NewEncoder(file)

	// Encode the BST items into the file using the json encoder
	err = encoder.Encode(items)
	if (err != nil) {
		return false
	}

	return true
}

// Deserialize is a public method that desearlizes a tree from the fiven file
// into the given tree, and returns true if successful and false if not.
func (tree *Tree) Deserialize(filename string) bool {
	// Check if filename is an empty string
	if (filename == "") {
		return false
	}

	// Open the input file
	file, err := os.Open(filename)
	if (err != nil) {
		return false
	}
	defer file.Close()

	// Create new decoder for the input file
	decoder := json.NewDecoder(file)

	var items []int

	// Decode the JSON representation of the tree from the input file
	err = decoder.Decode(&items)
	if (err != nil) {
		return false
	}

	// Reset the tree to the initial state
	tree.root = nil
	tree.nodes = 0

	// Insert each item into the tree
	for _, v := range items {
		tree.Insert(v)
	}

	return true
}