package dev04

import (
	"maps"
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	example := []string{
		"пятка", "пятак", "Пятак", "тяпка", "листок", "стОЛик", "слиток", "абоба",
	}
	expected := map[string][]string{
		"пятка":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}

	actual := findAnagrams(example)
	if !maps.EqualFunc(actual, expected, func(s1 []string, s2 []string) bool {
		return slices.Compare(s1, s2) == 0
	}) {
		t.Fail()
	}
}
