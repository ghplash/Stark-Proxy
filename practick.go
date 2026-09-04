package main

import "fmt"





func doing() {
	list := &DoubleLinkedList{}
	list.PushFront("Google.com")
	list.PushFront("Yandex.ru")
	list.PushFront("GitHub.com")
	list.printList()
	list.RemoveTail()
	list.printList()
}





func (D *DoubleLinkedList) printList() {
	current := D.Head
	for current != nil {
		fmt.Print(current.Valye, "<--->")
		current = current.Next
	}
	fmt.Println()
}
