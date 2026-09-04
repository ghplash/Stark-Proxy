package main

type Node struct {
	Valye string
	Next  *Node
	Prev  *Node
}

type DoubleLinkedList struct {
	Head *Node
	Tail *Node
}

type CachePool struct {
	Data map[string]*Node
	List *DoubleLinkedList
}

func NewCachePool() *CachePool {
	return &CachePool{
		Data: make(map[string]*Node),
		List: &DoubleLinkedList{},
	}
}

func (D *DoubleLinkedList) PushFront(valye string) {
	newNode := &Node{Valye: valye}
	if D.Head == nil {
		D.Head = newNode
		D.Tail = newNode
	} else {
		newNode.Next = D.Head
		D.Head.Prev = newNode
		D.Head = newNode
	}
}

func (D *DoubleLinkedList) RemoveTail() {
	if D.Tail == nil {
		return
	}
	if D.Tail == D.Head {
		D.Head = nil
		D.Tail = nil
	} else {
		D.Tail = D.Tail.Prev
		D.Tail.Next = nil
	}
}
