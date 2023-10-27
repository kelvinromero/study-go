package two_sum

// Time: O(n²)
//func twoSum(nums []int, target int) []int {
//	var result []int
//	for i, num := range nums {
//		for j, num2 := range nums {
//			if i != j {
//				if num+num2 == target {
//					result = append(result, i, j)
//					return result
//				}
//			}
//		}
//	}
//	return result
//}

// Time: O(n)
// Space: O(n)
func twoSum(nums []int, target int) []int {
	var numMap = make(map[int]int)

	for i, num := range nums {
		pos, found := numMap[target-num]

		if found && pos != i {
			return []int{pos, i}
		}

		numMap[num] = i
	}

	return []int{}
}
