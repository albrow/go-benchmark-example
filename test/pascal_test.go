package test

import (
	"github.com/albrow/go-pascal/common"
	"github.com/albrow/go-pascal/implementations"
	"testing"
)

func TestNaive(t *testing.T) {
	testPascaler(t, implementations.Naive)
}

func testPascaler(t *testing.T, p common.Pascaler) {
	// cases is an array of test cases, each consisting of two inputs
	// n and m and the expected output.
	cases := []struct {
		n        int
		m        int
		expected int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{4, 1, 4},
		{6, 4, 15},
		{7, 3, 35},
		{7, 0, 1},
		{7, 7, 1},
	}

	// Iterate through each test case and check the result
	for _, c := range cases {
		got := p.Pascal(c.n, c.m)
		if got != c.expected {
			t.Errorf("Incorrect result for %s with inputs (%d, %d).\nExpected %d but got %d.", p, c.n, c.m, c.expected, got)
		}
	}
}