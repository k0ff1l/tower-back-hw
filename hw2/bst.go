package main

import "fmt"

// binary search tree

type NodeBST struct {
	val   int
	left  *NodeBST
	right *NodeBST
}

type BST struct {
	root *NodeBST
}

func NewTree() *BST {
	return &BST{root: nil}
}

func PrintFromNode(node *NodeBST, level int) {
	if node == nil {
		return
	}

	PrintFromNode(node.right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("	") // tabulation
	}
	fmt.Println(node.val)
	PrintFromNode(node.left, level+1)
}

func (tree *BST) AddNode(val int) {
	tree.root = addNode(val, tree.root)
}

func addNode(val int, node *NodeBST) *NodeBST {
	if node == nil {
		return &NodeBST{val: val}
	}

	if node.val == val {
		return node
	}

	if val < node.val {
		node.left = addNode(val, node.left)
	} else {
		node.right = addNode(val, node.right)
	}

	return node
}

func (tree *BST) DeleteNode(val int) {
	if tree.root != nil {
		tree.root = deleteNode(val, tree.root)
	}
}

func deleteNode(val int, node *NodeBST) *NodeBST {
	if node == nil {
		return nil
	}

	if val < node.val {
		node.left = deleteNode(val, node.left)
	} else if val > node.val {
		node.right = deleteNode(val, node.right)
	} else {
		// todo: проверить когда удаляем обрубок (лист)
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minNode := minValueNode(node.right)
		node.val = minNode.val
		node.right = deleteNode(minNode.val, node.right)
	}

	return node
}

func minValueNode(node *NodeBST) *NodeBST {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (tree *BST) IsExist(val int) bool {
	return isExist(val, tree.root)
}

func isExist(val int, node *NodeBST) bool {
	if node == nil {
		return false
	}

	if val < node.val {
		return isExist(val, node.left)
	} else if val > node.val {
		return isExist(val, node.right)
	}

	return true
}

func TreeTest() {
	bst := NewTree()

	// Test 1: Add multiple nodes
	bst.AddNode(10)
	bst.AddNode(5)
	bst.AddNode(15)
	bst.AddNode(3)
	bst.AddNode(7)
	bst.AddNode(12)
	bst.AddNode(18)

	fmt.Println("BST values after adding nodes:")
	PrintFromNode(bst.root, 0)

	// Test 2: Delete a leaf node
	bst.DeleteNode(3)
	fmt.Println("\nBST values after deleting leaf node 3:")
	PrintFromNode(bst.root, 0)

	// Test 3: Delete a node with one child
	bst.DeleteNode(7)
	fmt.Println("\nBST values after deleting node 7 (one child):")
	PrintFromNode(bst.root, 0)

	// Test 4: Delete a node with two children
	bst.DeleteNode(15)
	fmt.Println("\nBST values after deleting node 15 (two children):")
	PrintFromNode(bst.root, 0)

	// Test 5: Delete root node
	bst.DeleteNode(10)
	fmt.Println("\nBST values after deleting the root node 10:")
	PrintFromNode(bst.root, 0)

	// Test 6: Deleting non-existent value
	bst.DeleteNode(100)
	fmt.Println("\nAttempting to delete non-existent value 100 (no change should happen):")
	PrintFromNode(bst.root, 0)

	// Test 7: Inserting duplicate values (should not be added)
	bst.AddNode(10)
	bst.AddNode(5) // duplicate
	fmt.Println("\nBST values after attempting to add duplicate 5:")
	PrintFromNode(bst.root, 0)

	if !bst.IsExist(5) {
		fmt.Print("failure")
	}

	if !bst.IsExist(18) {
		fmt.Print("failure")
	}

	if !bst.IsExist(10) {
		fmt.Print("failure")
	}

	if !bst.IsExist(12) {
		fmt.Print("failure")
	}

	if bst.IsExist(100) {
		fmt.Print("failure")
	}

}
