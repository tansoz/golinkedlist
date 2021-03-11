package golinkedlist

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestLinkedList_String(t *testing.T) {

	a := []int{1, 2, 3, 4, 5, 6}
	//a := []string{"aabc","1asdas11","accb","abc","bcc","abc","asdas115","cba","abbc","d","asdasd"}
	ll := NewLinkedList()

	for i := range a {
		ll.HeadInsert(unsafe.Pointer(&a[i]))
	}

	//ll.TailInsert(unsafe.Pointer(&a[5]))

	ll.ForEach(func(node *Node, index int64) bool {
		fmt.Println(*(*int)(node.data))
		return false
	})

	/*fmt.Println(ll)

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
	})*/
	fmt.Println("---------------------------")
	/*ll.RangeSort(func(a unsafe.Pointer, b unsafe.Pointer) int {
		fmt.Println(*(*int)(a), *(*int)(b))
		return *(*int)(b) - *(*int)(a)
	}, nil, nil)*/
	ll.RangeSort(nil, nil, nil)
	fmt.Println("---------------------------")

	ll.ForEach(func(node *Node, index int64) bool {
		fmt.Println("RESULT:", *(*int)(node.data))
		return true
	})

	for i := range ll.Items() {
		fmt.Println(*(*int)(i.GetData()))
	}

	fmt.Println("-----------------------------")

	ll.Shuffle(0)

	for i := range ll.Items() {
		fmt.Println(*(*int)(i.GetData()))
	}

}

func TestThreadSafe(t *testing.T) {

	ll := NewLinkedList()
	a := ""
	ll.HeadInsert(unsafe.Pointer(&a))
	go func() {

		for {
			a := time.Now().String()
			ll.GetHeadNode().SetData(unsafe.Pointer(&a))
		}

	}()

	for {

		fmt.Println(*(*string)(ll.GetHeadNode().GetData()))

	}

}
