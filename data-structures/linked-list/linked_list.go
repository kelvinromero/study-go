package linked_list

import (
	"fmt"
	"io"
)

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Append(data int) {
	newNode := Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = &newNode
		return
	}

	node := ll.Head
	for node.Next != nil {
		node = node.Next
	}

	node.Next = &newNode
}

func (ll *LinkedList) Display(writer io.Writer) {
	node := ll.Head
	for node != nil {
		if node.Next == nil {
			fmt.Fprintf(writer, "%d", node.Data)
			break
		}

		fmt.Fprintf(writer, "%d -> ", node.Data)
		node = node.Next
	}
}
