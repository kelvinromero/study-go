package group_anagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	checkAnagrams := func(t testing.TB, got, want [][]string) {
		t.Helper()

		if !compareAnagrams(got, want) {
			t.Errorf("\n got %v, \n want %v", got, want)
		}
	}

	t.Run("two anagrams", func(t *testing.T) {
		got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
		want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
		checkAnagrams(t, got, want)
	})

	t.Run("two anagrams", func(t *testing.T) {
		got := groupAnagrams([]string{""})
		want := [][]string{{""}}
		checkAnagrams(t, got, want)
	})

	t.Run("two anagrams", func(t *testing.T) {
		got := groupAnagrams([]string{"a"})
		want := [][]string{{"a"}}
		checkAnagrams(t, got, want)
	})

	t.Run("two anagrams", func(t *testing.T) {
		got := groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"})
		want := [][]string{{"cab"}, {"tin"}, {"pew"}, {"duh"}, {"may"}, {"ill"}, {"buy"}, {"bar"}, {"max"}, {"doc"}}
		checkAnagrams(t, got, want)
	})
}

func compareAnagrams(got [][]string, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}

	for i, v := range got {
		if len(v) != len(want[i]) {
			return false
		}
	}

	return true
}
