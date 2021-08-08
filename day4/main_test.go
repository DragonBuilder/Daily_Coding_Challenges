package main

import (
	"testing"
)

type testcase struct {
	input    []int
	expected int
}

var testcases = []testcase{
	{input: []int{3, 4, -1, 1}, expected: 2},
	{input: []int{1, 2, 0}, expected: 3},
}

func TestFindLowestPositive(t *testing.T) {
	for _, tc := range testcases {
		got := findLowestPositive(tc.input)
		if got != tc.expected {
			t.Errorf("findMissing(%v) = %d; want = %d", tc.input, got, tc.expected)
		}
	}
}

func TestSort(t *testing.T) {
	L := []int{3, 4, -1, 1}

	expected := []int{-1, 1, 3, 4}
	sort(L)
	for i := range L {
		if L[i] != expected[i] {
			t.Errorf("expected %v got %v", expected, L)
		}
	}
}
