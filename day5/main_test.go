package main

import "testing"

type testcase struct {
	a        int
	b        int
	expected int
}

var carTestcases = []testcase{
	{a: 1, b: 2, expected: 1},
	{a: -3, b: 0, expected: -3},
	{a: 9, b: 9, expected: 9},
}

var cdrTestcases = []testcase{
	{a: 1, b: 2, expected: 2},
	{a: -3, b: 0, expected: 0},
	{a: 9, b: 9, expected: 9},
}

func TestCar(t *testing.T) {
	for _, tc := range carTestcases {
		got := car(cons(tc.a, tc.b))
		if got != tc.expected {
			t.Errorf("car(cons(%d, %d)) = %d; want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}

func TestCdr(t *testing.T) {
	for _, tc := range cdrTestcases {
		got := cdr(cons(tc.a, tc.b))
		if got != tc.expected {
			t.Errorf("cdr(cons(%d, %d)) = %d; want %d", tc.a, tc.b, got, tc.expected)
		}
	}
}
