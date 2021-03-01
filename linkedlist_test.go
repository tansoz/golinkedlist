package golinkedlist

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestLinkedList_String(t *testing.T) {

	a := []int{6, 5, 4, 2, 1, 8, 7, 3, 9}
	ll := NewLinkedList()

	for i := range a {
		ll.HeadInsert(unsafe.Pointer(&a[i]))
	}

	fmt.Println(ll)

	ll.Reverse()

	fmt.Println(ll)

	ll.ForEach(func(node *Node) bool {
		fmt.Println(*(*int)(node.data))
		return false
	})

	ll.Move(ll.GetHeadNode(), ll.GetTailNode())

	ll.ForEach(func(node *Node) bool {
		fmt.Println(*(*int)(node.data))
		return false
	})

	ll.RangeSort(func(a unsafe.Pointer, b unsafe.Pointer) int {
		fmt.Println(*(*int)(a), *(*int)(b))
		return *(*int)(a) - *(*int)(b)
		//return -1
	}, ll.head, ll.tail)
	fmt.Println("---------------------------")
	ll.ForEach(func(node *Node) bool {
		fmt.Println(*(*int)(node.data))
		return false
	})

}
