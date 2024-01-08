package dev02

import "testing"

type response struct {
	data string
	err  error
}

func TestDecodeRepeats(t *testing.T) {
	examples := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"a12bc4d6e5",
	}
	expected := []response{
		{"aaaabccddddde", nil},
		{"abcd", nil},
		{"", ErrSyntax},
		{"", nil},
		{"aaaaaaaaaaaabccccddddddeeeee", nil},
	}

	for i, ex := range examples {
		actual := response{}
		actual.data, actual.err = DecodeRepeats(ex)
		if actual != expected[i] {
			t.Errorf(
				"answer wrong, testcase: %s\nexpected:%v\nactual:%v\n",
				examples[i],
				expected[i],
				actual,
			)
		}
	}
}
