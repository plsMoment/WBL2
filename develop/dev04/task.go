package dev04

import (
	"slices"
	"strings"
)

// helper returns map with key sorted anagram and value anagrams
func helper(s []string) map[string][]string {
	res := make(map[string][]string)
	for _, val := range s {
		val = strings.ToLower(val)
		runeStr := []rune(val)
		slices.Sort(runeStr)
		keyStr := string(runeStr)
		res[keyStr] = append(res[keyStr], val)
	}
	return res
}

// findAnagrams returns correct map
func findAnagrams(s []string) map[string][]string {
	withFakeKey := helper(s)
	withRightKey := make(map[string][]string, len(withFakeKey))

	for _, val := range withFakeKey {
		withRightKey[val[0]] = val
	}

	for key, val := range withRightKey {
		slices.Sort(val)
		withRightKey[key] = slices.Compact(val)
		if len(withRightKey[key]) < 2 {
			delete(withRightKey, key)
		}
	}
	return withRightKey
}
