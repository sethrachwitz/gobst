package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := NewBST()

	// Test normal inserts
	values := []int{4, 3, 6, 2, 7, 5}

	for _, v := range values {
		if (tree.Insert(v) == false) {
			t.Error("Node not inserted correctly")
		}
	}

	// Test insert of existing value (root)
	if (tree.Insert(4) == true) {
		t.Error("Existing node (root) inserted again")
	}

	// Test insert of existing value (non-root, non-leaf)
	if (tree.Insert(3) == true) {
		t.Error("Existing node (non-root, non-leaf) inserted again")
	}

	// Test insert of existing value (leaf)
	if (tree.Insert(7) == true) {
		t.Error("Existing node (leaf) inserted again")
	}

	// Test that tree structure is correct
	if (tree.root.value != 4 || 
		tree.root.left.value != 3 || 
		tree.root.right.value != 6 && 
		tree.root.left.left.value != 2 || 
		tree.root.right.left.value != 5 || 
		tree.root.right.right.value != 7) {
			t.Error("Tree structure is not correct")
	}

	// Test that node count is correct
	if (tree.nodes != 6) {
		t.Error("Tree node count is incorrect")
	}
}

func TestFind(t *testing.T) {
	// Setup code, test fails if setup fails
	tree := NewBST()

	values := []int{4, 3, 6, 2, 7, 5}

	for _, v := range values {
		if (tree.Insert(v) == false) {
			t.Fatal("Setup failed")
		}
	}

	// Test that tree structure is correct
	if (tree.root.value != 4 || 
		tree.root.left.value != 3 || 
		tree.root.right.value != 6 && 
		tree.root.left.left.value != 2 || 
		tree.root.right.left.value != 5 || 
		tree.root.right.right.value != 7) {
			t.Fatal("Setup failed")
	}

	// Test that node count is correct
	if (tree.nodes != 6) {
		t.Fatal("Setup failed")
	}

	// Run tests

	// Find root node
	found, level := tree.Find(4)
	if found != true {
		t.Error("Root node 4 not found")
	} else if level != 1 {
		t.Error("Root node 4 found at incorrect level")
	}

	// Find node requiring traversal left and right to find
	found, level = tree.Find(5)
	if found != true {
		t.Error("Node 5 not found")
	} else if level != 3 {
		t.Error("Node 5 found at incorrect level")
	}
}