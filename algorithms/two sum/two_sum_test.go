package two_sum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !compareSlices(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("two sum", func(t *testing.T) {
		got := twoSum([]int{2, 7, 11, 15}, 9)
		want := []int{0, 1}
		checkSums(t, got, want)
	})

	t.Run("two sum", func(t *testing.T) {
		got := twoSum([]int{3, 2, 4}, 6)
		want := []int{1, 2}
		checkSums(t, got, want)
	})

	t.Run("two sum", func(t *testing.T) {
		got := twoSum([]int{3, 3}, 6)
		want := []int{0, 1}
		checkSums(t, got, want)
	})

	t.Run("two sum", func(t *testing.T) {
		got := twoSum([]int{3, 2, 3}, 6)
		want := []int{0, 2}
		checkSums(t, got, want)
	})

	t.Run("two sum", func(t *testing.T) {
		var nums []int
		for i := 1; i <= 10000; i++ {
			nums = append(nums, i)
		}
		got := twoSum(nums, 19999)
		want := []int{9998, 9999}
		checkSums(t, got, want)
	})
}

func compareSlices(got []int, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i, v := range got {
		if v != want[i] {
			return false
		}
	}

	return true
}
