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

    - func (linkedList *LinkedList) ForEach(fn func(node *Node) bool)
        ```go
        linkedList.ForEach(func(node *Node, index int64)bool{
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
        > parameter fn can use nil
        ```go
        // 5 6 2 1 3 4
        linkedList.RangeSort(function(a unsafe.Pointer, b unsafe.Pointer){
            return *(int)(b) - *(int)(a)
        }int,nil,nil)
        // 1 2 3 4 5 6
        linkedList.RangeSort(function(a unsafe.Pointer, b unsafe.Pointer){
            return *(int)(a) - *(int)(b)
        }int,nil,nil)
        // 6 5 4 3 2 1
        linkedList.RangeSort(nil,nil,nil)
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
    - func (linkedList *LinkedList) Get(index int64, data unsafe.Pointer) unsafe.Pointer
        ```go
        node := linkedList.Get(5)
        ```
    - func (linkedList *LinkedList) Set(index int64, data unsafe.Pointer) unsafe.Pointer
        ```go
        a := 1
        linkedList.Set(0,unsafe.Pointer(&a))
        ```
    - func (linkedList *LinkedList) GetCyclePrevNode(node *Node) *Node
        ```go
        linkedList.GetCyclePrevNode(node)
        ```
    - func (linkedList *LinkedList) GetCycleNextNode(node *Node) *Node
        ```go
        linkedList.GetCycleNextNode(node)
        ```
    - func (linkedList *LinkedList) ReplaceNode(replaceNode *Node, targetNode *Node) *Node
        ```go
        linkedList.ReplaceNode(node1,node2)
        ```
    - func (linkedList *LinkedList) Append(data unsafe.Pointer)
        ```go
        // 1->2->3
        linkedList.Append(data) // Assuming the value of data is 4
        // 1->2->3->4
        ```
    - func (linkedList *LinkedList) Prepend(data unsafe.Pointer)
        ```go
        // 1->2->3
        linkedList.Prepend(data)    // Assuming the value of data is 4
        // 4->1->2->3
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