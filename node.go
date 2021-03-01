package golinkedlist

import (
	"fmt"
	"unsafe"
)

type Node struct {
	next *Node
	prev *Node
	data unsafe.Pointer // Pointer to data address
}

/* New a linked-list node */
func NewNode(data unsafe.Pointer) *Node {
	return &Node{
		next: nil,
		prev: nil,
		data: data,
	}
}

func (node *Node) setNextNode(nextNode *Node) *Node {
	tmp := node.next
	node.next = nextNode
	return tmp
}

func (node *Node) setPrevNode(prevNode *Node) *Node {
	tmp := node.prev
	node.prev = prevNode
	return tmp
}

func (node *Node) SetData(data unsafe.Pointer) unsafe.Pointer {
	tmp := node.data
	node.data = data
	return tmp
}

func (node *Node) GetNextNode() *Node {
	return node.next
}

func (node *Node) GetPrevNode() *Node {
	return node.prev
}

func (node *Node) GetData() unsafe.Pointer {
	return node.data
}

/* Swap two nodes position */
func (node *Node) Swap(b *Node) {

	var (
		tmpNode *Node
	)
	if b != nil {
		tmpNode = node.GetNextNode()
		node.setNextNode(b.GetNextNode())
		b.setNextNode(tmpNode)

		tmpNode = node.GetPrevNode()
		node.setPrevNode(b.GetPrevNode())
		b.setPrevNode(tmpNode)

		if node.GetPrevNode() != nil {
			node.GetPrevNode().setNextNode(node)
		}
		if node.GetNextNode() != nil {
			node.GetNextNode().setPrevNode(node)
		}
		if b.GetNextNode() != nil {
			b.GetNextNode().setPrevNode(b)
		}

	}
}

/* The function fmt.Println default call this function to print the string */
func (node *Node) String() string {
	return fmt.Sprintf("Node(0x%x){data:0x%x}", unsafe.Pointer(node), node.data)
}
