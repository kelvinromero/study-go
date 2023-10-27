package group_anagrams

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) (result [][]string) {
	if len(strs) == 0 {
		return
	}

	var strMap = make(map[string][]string)

	// Time: O(n)
	for _, s := range strs {
		ss := sortedString(s)
		strMap[ss] = append(strMap[ss], s)
	}

	for _, arr := range strMap {
		result = append(result, arr)
	}

	return
}

func sortedString(s string) string {
	x := strings.Split(s, "")
	slices.Sort(x)
	s = strings.Join(x, "")

	return s
}
