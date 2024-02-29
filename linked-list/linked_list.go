package linked_list 

import (
	"errors"
	"fmt"
)

type List struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func MainLinkedList() {
	l := NewList()
	l.InsertLast(20)
	l.InsertLast(30)
	l.InsertLast(40)
	l.InsertFirst(50)

	l.Contains(40)
	l.Contains(60)
	l.PrintList()
}

func NewList() *List {
	return &List{
		Head: nil,
	}
}

func (l *List) InsertFirst(value int) {
	insertNode := &Node{
		Value: value,
	}

	insertNode.Next = l.Head
	l.Head = insertNode
}

func (l *List) InsertLast(value int) {
	insertNode := &Node{
		Value: value,
	}

	if l.Head == nil {
		l.Head = insertNode
		return
	}

	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = insertNode
}

func (l *List) Remove()

func (l *List) IsEmpty() bool {
	return l.Head == nil
}

func (l *List) GetFirst() *Node {
	return l.Head
}

func (l *List) GetLast() *Node {
	curr := l.Head

	if !l.IsEmpty() {
		for curr.Next != nil {
			curr = curr.Next
		}
	}

	return curr
}

func (l *List) GetIndex(i int) (*Node, error) {
    if l.Size() < i || i < 0 {
        return nil, errors.New("The index must be between 0 and the list size")
    }

    if l.IsEmpty() {
        return nil, errors.New("The list is empty")
    }
    
    curr := l.Head
    count := 0
    // TODO - Analyze this condition
    for curr != nil {
        count++
        if count == i {
            break
        }

        curr = curr.Next
    }

    return curr, nil
}

func (l *List) Size() (int) {
    size := 0
    curr := l.Head

    if !l.IsEmpty() {
        for curr.Next != nil {
            size++
            curr = curr.Next
        }
    }

    return size
}

func (l *List) Contains(value int) {
	curr := l.Head
	for curr != nil {
		if curr.Value == value {
			fmt.Println(true)
			return
		}

		curr = curr.Next
	}

	fmt.Println(false)
	return
}

func (l *List) PrintList() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d\n", curr.Value)
		curr = curr.Next
	}
}
