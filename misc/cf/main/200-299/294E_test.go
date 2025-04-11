//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/294/E
// https://codeforces.com/problemset/status/294/problem/E
func TestCF294E(t *testing.T) {
	t.Log("Current test is [CF294E]")
	testCases := [][2]string{
		{
			`3
			1 2 2
			1 3 4`,
			`12`,
		},
		{
			`6
			1 2 1
			2 3 1
			3 4 1
			4 5 1
			5 6 1`,
			`29`,
		},
		{
			`6
			1 3 1
			2 3 1
			3 4 100
			4 5 2
			4 6 1`,
			`825`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF294E, testCases, 0)
}
