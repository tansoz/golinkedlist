package golinkedlist

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewNode(t *testing.T) {

	data := "I'm the NO.1."
	node := NewNode(unsafe.Pointer(&data))
	fmt.Println(node)

	fmt.Println((*string)(node.GetData()))

}
