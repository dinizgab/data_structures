package bst

import "fmt"

type BinaryTree struct {
	Head *TreeNode
}

type TreeNode struct {
	Right *TreeNode
	Left  *TreeNode
	Value int
}

func MainBST() {
	tree := NewBST()

	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(12)
	tree.Insert(9)
	tree.Insert(16)

    tree.PreOrder(tree.Head)
}

func NewBST() *BinaryTree {
	return &BinaryTree{
		Head: nil,
	}
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.Head == nil
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Head == nil {
		tree.Head = &TreeNode{Value: value}
	} else {
		tree.Head.insert(value)
	}
}

func (node *TreeNode) insert(value int) {
	if value > node.Value {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value}
		} else {
			node.Right.insert(value)
		}
	} else {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value}
		} else {
			node.Left.insert(value)
		}
	}
}

func (tree *BinaryTree) PreOrder(curr *TreeNode) {
	if curr != nil {
		fmt.Printf("%d ", curr.Value)
		tree.PreOrder(curr.Left)
		tree.PreOrder(curr.Right)
	}
}
