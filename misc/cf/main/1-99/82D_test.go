//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/82/D
// https://codeforces.com/problemset/status/82/problem/D
func TestCF82D(t *testing.T) {
	t.Log("Current test is [CF82D]")
	testCases := [][2]string{
		{
			`4
			1 2 3 4`,
			`6
			1 2
			3 4`,
		},
		{
			`5
			2 4 3 1 4`,
			`8
			1 3
			2 5
			4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF82D, testCases, 0)
}
