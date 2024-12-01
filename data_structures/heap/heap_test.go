package heap

import "testing"

func TestHeapSize(t *testing.T) {
	heap := NewHeap(10)
	if heap.Size() != 0 {
		t.Errorf("heap size should be 0")
	}
}
