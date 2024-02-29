package bst

import (
	"fmt"
)

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

    tree.Search(12)
    tree.Search(8)

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

func (tree *BinaryTree) Search(value int) {
	if tree.IsEmpty() {
		fmt.Println("The list is empty!")
		return 
	}

    fmt.Println(tree.search(value, tree.Head))
}

func (tree *BinaryTree) search(value int, curr *TreeNode) (int) {
    if curr == nil {
        return -1
    } else if value == curr.Value {
		return curr.Value
	} else if value > curr.Value {
		return tree.search(value, curr.Right)
	} else {
		return tree.search(value, curr.Left)
	}
}

func (tree *BinaryTree) PreOrder(curr *TreeNode) {
	if curr != nil {
		fmt.Printf("%d ", curr.Value)
		tree.PreOrder(curr.Left)
		tree.PreOrder(curr.Right)
	}
}
