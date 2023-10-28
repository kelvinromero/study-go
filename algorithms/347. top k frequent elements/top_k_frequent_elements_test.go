package top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func TestTopKFrequentElements(t *testing.T) {

	testName := "test top k frequenet elements"
	elementTests := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{nums: []int{1, 1}, k: 1, want: []int{1}},
	}

	for _, tt := range elementTests {
		t.Run(testName, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v got %v want %v", testName, got, tt.want)
			}
		})
	}
}
