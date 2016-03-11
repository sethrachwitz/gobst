package bst

import (
	"testing"
)

func SetupBST() (tree *Tree) {
	tree = NewBST()

	values := []int{4, 3, 6, 2, 7, 5}

	for _, v := range values {
		if (tree.Insert(v) == false) {
			return nil
		}
	}

	// Test that tree structure is correct
	if (tree.root.value != 4 || 
		tree.root.left.value != 3 || 
		tree.root.right.value != 6 && 
		tree.root.left.left.value != 2 || 
		tree.root.right.left.value != 5 || 
		tree.root.right.right.value != 7) {
			return nil
	}

	// Test that node count is correct
	if (tree.nodes != 6) {
		return nil
	}

	return
}

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
	tree := SetupBST()

	if (tree == nil) {
		t.Fatal("Setup failed")
	}

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

func TestDelete(t *testing.T) {
	tree := SetupBST()

	if (tree == nil) {
		t.Fatal("Setup failed")
	}

	// Delete root node
	if (!tree.Delete(4)) {
		t.Error("Could not delete root node")
	}

	// Delete non-root and non-leaf node
	if (!tree.Delete(3)) {
		t.Error("Could not delete non-root and non-leaf node")
	}

	// Delete leaf node
	if (!tree.Delete(7)) {
		t.Error("Could not delete leaf node")
	}

	// Check that tree structure is okay
	correct := []int{2,5,6}

	items := tree.InOrder()

	if (len(items) != len(correct)) {
		t.Error("Tree structure is incorrect")
	} else {
		for k, v := range items {
			if v != correct[k] {
				t.Error("Tree structure is incorrect")
			}
		}
	}
}

func TestPreOrder(t *testing.T) {
	tree := SetupBST()

	if (tree == nil) {
		t.Fatal("Setup failed")
	}

	items := tree.PreOrder()

	correct := []int{4,3,2,6,5,7}

	for k, v := range items {
		if (v != correct[k]) {
			t.Error("PreOrder item at index", k, "is incorrect")
		}
	}
}

func TestInOrder(t *testing.T) {
	tree := SetupBST()

	if (tree == nil) {
		t.Fatal("Setup failed")
	}

	items := tree.InOrder()

	correct := []int{2,3,4,5,6,7}

	for k, v := range items {
		if (v != correct[k]) {
			t.Error("InOrder item at index", k, "is incorrect")
		}
	}
}

func TestPostOrder(t *testing.T) {
	tree := SetupBST()

	if (tree == nil) {
		t.Fatal("Setup failed")
	}

	items := tree.PostOrder()

	correct := []int{2,3,5,7,6,4}

	for k, v := range items {
		if (v != correct[k]) {
			t.Error("PostOrder item at index", k, "is incorrect")
		}
	}
}