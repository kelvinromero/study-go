package linked_list

import (
	"bytes"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("test display", func(t *testing.T) {
		buffer := bytes.Buffer{}

		linkedList := LinkedList{}
		linkedList.Append(1)
		linkedList.Append(2)
		linkedList.Append(3)
		linkedList.Append(4)

		linkedList.Display(&buffer)

		got := buffer.String()
		want := "1 -> 2 -> 3 -> 4"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
