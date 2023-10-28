package top_k_frequent_elements

import (
	"sort"
)

func topKFrequent(nums []int, k int) (result []int) {
	numMap := make(map[int]int)
	var sortedNCs []NumCount

	// O(n)
	for _, n := range nums {
		numMap[n] += 1
	}

	// O(n)
	for num, count := range numMap {
		println(num, count)
		sortedNCs = append(sortedNCs, NumCount{num, count})
	}

	sort.Slice(sortedNCs, func(i, j int) bool {
		return sortedNCs[i].count > sortedNCs[j].count
	})

	for _, nc := range sortedNCs[:k] {
		result = append(result, nc.num)
	}

	return
}

type NumCount struct {
	num   int
	count int
}
