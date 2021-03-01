package golinkedlist

import (
	"fmt"
	"unsafe"
)

type LinkedList struct {
	head *Node // The head node of linked-list
	tail *Node // The tail node of linked-list
}

/* Init a linked-list */
func NewLinkedList() *LinkedList {

	return &LinkedList{
		head: nil,
		tail: nil,
	}
}

func (linkedList *LinkedList) GetHeadNode() *Node {
	return linkedList.head
}

func (linkedList *LinkedList) GetTailNode() *Node {
	return linkedList.tail
}

/* Insert node into thr linked-list */
func (linkedList *LinkedList) InsertNode(prevNode *Node, nextNode *Node, node *Node) {

	node.setPrevNode(prevNode)
	node.setNextNode(nextNode)

	if linkedList.head == nil || linkedList.tail == nil {
		linkedList.head = node
		linkedList.tail = node
	} else if nextNode == linkedList.head {
		linkedList.head.setPrevNode(node)
		linkedList.head = node
	} else if prevNode == linkedList.tail {
		linkedList.tail.setNextNode(node)
		linkedList.tail = node
	} else {
		prevNode.next = node
		nextNode.prev = node
	}
}

/* Insert the data to the linked-list */
func (linkedList *LinkedList) Insert(prevNode *Node, nextNode *Node, data unsafe.Pointer) *Node {

	node := NewNode(data) // Initial a new node of the linked-list

	linkedList.InsertNode(prevNode, nextNode, node)

	return node
}

/* New insert will be the linked-list head */
func (linkedList *LinkedList) HeadInsert(data unsafe.Pointer) {
	linkedList.Insert(nil, linkedList.head, data)
}

/* New insert will be the linked-list tail */
func (linkedList *LinkedList) TailInsert(data unsafe.Pointer) {
	linkedList.Insert(linkedList.tail, nil, data)
}

/* Remove the node of the linked-list */
func (linkedList *LinkedList) RemoveNode(node *Node) {

	if node == linkedList.head {
		linkedList.head = node.GetNextNode()
	}
	if node == linkedList.tail {
		linkedList.tail = node.GetPrevNode()
	}
	if node.GetPrevNode() != nil {
		node.GetPrevNode().setNextNode(node.GetNextNode())
	}
	if node.GetNextNode() != nil {
		node.GetNextNode().setPrevNode(node.GetPrevNode())
	}
}

/* Foreach all node of the linked-list */
func (linkedList *LinkedList) ForEach(fn func(node *Node) bool) *Node {

	for currentNode := linkedList.head; currentNode != nil; currentNode = currentNode.GetNextNode() {
		if fn(currentNode) {
			return currentNode
		}
	}
	return nil
}

/* Reverse the linked-list node */
func (linkedList *LinkedList) Reverse() {

	var currNode *Node

	for currNode = linkedList.head; currNode != nil; {
		currNode = currNode.setNextNode(currNode.setPrevNode(currNode.GetNextNode()))
	}

	currNode = linkedList.head
	linkedList.head = linkedList.tail
	linkedList.tail = currNode
}

/* Count the length of the linked-list */
func (linkedList *LinkedList) Len() int64 {

	count := int64(0)

	linkedList.ForEach(func(node *Node) bool {
		count += 1
		return false
	})
	return count
}

/* Swap two nodes position in the linked-list */
func (linkedList *LinkedList) Swap(a *Node, b *Node) (*Node, *Node) {

	if b == linkedList.head {
		linkedList.head = a
	}
	if a == linkedList.head {
		linkedList.head = b
	}
	if b == linkedList.tail {
		linkedList.tail = a
	}
	if a == linkedList.tail {
		linkedList.tail = b
	}

	a.Swap(b)
	return b, a
}

/* Move the node from a place to the front of the target node in linked-list */
func (linkedList *LinkedList) Move(node *Node, target *Node) {

	if node != nil && target != nil && node != target {

		linkedList.RemoveNode(node)                               // Remove the node from the linked-list
		linkedList.InsertNode(target.GetPrevNode(), target, node) // Then, insert it into the linked-list
	}
}

/* Show the linked-list */
func (linkedList *LinkedList) String() string {

	s := fmt.Sprintf("LinkedLst(%d)[\n", linkedList.Len())

	linkedList.ForEach(func(node *Node) bool {

		s += "\t" + node.String() + "\n"
		return false
	})

	return s + "]"
}

/* Sort the link */
/* Must be from the from node to the to node. continuously. */
/*

original: 3->1->5->2->4

1->3->5->2->4
2->1->3->5->4
1->2->3->5->4
1->2->3->4->4

*/
func (linkedList *LinkedList) RangeSort(fn func(a unsafe.Pointer, b unsafe.Pointer) int, from *Node, to *Node) {

	var (
		low   *Node
		high  *Node
		pivot *Node
		tmp   *Node
	)

	pivot = from
	low = from
	high = from

	if low != nil && high != nil && from != to {

		for high != nil {

			if high != pivot && fn(pivot.GetData(), high.GetData()) < 0 {
				tmp = high.GetPrevNode()
				linkedList.Move(high, low)
				low = high
				high = tmp
			}
			if high == to || low == to {
				break
			}
			high = high.GetNextNode()
		}

		//fmt.Println(*(*int)(low.GetData()),*(*int)(pivot.GetData()),*(*int)(high.GetData()))

		if pivot != low && pivot.GetPrevNode() != nil {
			linkedList.RangeSort(fn, low, pivot.GetPrevNode())
		}

		if pivot != high && pivot.GetNextNode() != nil {
			linkedList.RangeSort(fn, pivot.GetNextNode(), high)
		}
	}

}
