# GoLinkedList

Install:
```
go get -u github.com/tansoz/golinkedlist
```

APIs:

- LinkedList

    - func NewLinkedList() *LinkedList
        ```go
        linkedList := NewLinkedList()       // new a LinkedList object  
        ```

    - func (linkedList *LinkedList) ForEach(fn func(node *Node) bool) *Node
        ```go
        node := linkedList.ForEach(func(node *Node)bool{
            return [condition]
        })
        ```
      
    - func (linkedList *LinkedList) GetHeadNode() *Node
        ```go
        node := linkedList.GetHeadNode()
        ```

    - func (linkedList *LinkedList) GetTailNode() *Node
        ```go
        node := linkedList.GetTailNode()
        ```

    - func (linkedList *LinkedList) HeadInsert(data unsafe.Pointer)
        ```go
        // 1 2 3 4 5
        a := 6
        linkedList.HeadInsert(unsafe.Pointer(&a))
        // 6 1 2 3 4 5
        ```

    - func (linkedList *LinkedList) Insert(prevNode *Node, nextNode *Node, data unsafe.Pointer) *Node
        ```go
        a := 123
        linkedList.Insert(nil,linkedList.GetHeadNode(),unsafe.Pointer(&a))
        ```
  
    - func (linkedList *LinkedList) InsertNode(prevNode *Node, nextNode *Node, node *Node)
        ```go
        a := 123
        node := NewNode(&a)
        linkedList.InsertNode(nil,linkedList.GetHeadNode(),node)
        ```

    - func (linkedList *LinkedList) Items() chan *Node
        ```go
        for node := range linkedList.Items(){
            fmt.Println(node)
        }
        ```

    - func (linkedList *LinkedList) Len() int64
        ```go
        len := linkedList.Len() // count the length of the linked-list
        ```
    - func (linkedList *LinkedList) Merge(list *LinkedList)
        ```go
        l1.Merge(l2)
        ```
    - func (linkedList *LinkedList) Move(node *Node, target *Node)
        ```go
        linkedList.Move(node,targetNode)
        ```
    - func (linkedList *LinkedList) RangeSort(fn func(a unsafe.Pointer, b unsafe.Pointer) int, from *Node, to *Node)
        ```go
        // 5 6 2 1 3 4
        linkedList.RangeSort(function(a unsafe.Pointer, b unsafe.Pointer){
            return *(int)(a) - *(int)(b)
        }int,nil,nil)
        // 1 2 3 4 5 6
        ```
    - func (linkedList *LinkedList) RemoveNode(node *Node)
        ```go
        linkedList.RemoveNode(node)
        ```
    - func (linkedList *LinkedList) Reverse()
        ```go
        // 5 2 3 1 4
        linkedList.Reverse()
        // 4 1 3 2 5
        ```
    - func (linkedList *LinkedList) String() string
        ```go
        linkedList.String()
        ```
    - func (linkedList *LinkedList) Swap(a *Node, b *Node) (*Node, *Node)
        ```go
        linkedList.Swap(node1,node2)
        ```
    - func (linkedList *LinkedList) TailInsert(data unsafe.Pointer)
        ```go
        // 1 2 3 4 5
        a := 6
        linkedList.TailInsert(unsafe.Pointer(&a))
        // 1 2 3 4 5 6
        ```
    - func (linkedList *LinkedList) ToArray() []unsafe.Pointer
        ```go
        linkedList.ToArray()
        ```
- Node

    - func NewNode(data unsafe.Pointer) *Node
        ```go
        a := 1
        node := NewNode(unsafe.Pointer(&a))
        ```
    - func (node *Node) GetData() unsafe.Pointer
        ```go
        a := 1
        node := NewNode(unsafe.Pointer(&a))
        fmt.Println(*(int)(node.GetData()))
        // 1
        ```
    - func (node *Node) GetNextNode() *Node
        ```go
        node.GetNextNode()
        ```
    - func (node *Node) GetPrevNode() *Node
        ```go
        node.GetPrevNode()
        ```
    - func (node *Node) SetData(data unsafe.Pointer) unsafe.Pointer
        ```go
        c := 2
        node.SetData(unsafe.Pointer(&c))
        ```
    - func (node *Node) String() string
        ```go
        node.String()
        ```
    - func (node *Node) Swap(b *Node)
        ```go
        node.String(node2)
        ```
    - func (node *Node) Iterator(flag direct) chan *Node
        ```go
        // 1->2->3->4->5
      
        for n := range node.Iterator(FORWARD){   // Assuming node is the third node. That is 3.
            fmt.Println(*(int)(n.GetData()))
        }
        /*
        output:
        3
        4
        5
        */
      
        for n := range node.Iterator(BACKWARD){   // Assuming node is the third node. That is 3.
            fmt.Println(*(int)(n.GetData()))
        }
        /*
        output:
        3
        2
        1
        */
        ```

### The feature i want but do not achieve yet:

- Thread safe
- ~~Merge Linked-list~~
- ~~To []unsafe.Pointer~~

### Some functions wait for me to fixed:

- RangeSort