//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/650/D
// https://codeforces.com/problemset/status/650/problem/D
func TestCF650D(t *testing.T) {
	t.Log("Current test is [CF650D]")
	testCases := [][2]string{
		{
			`4 4
			1 2 3 4
			1 1
			1 4
			4 3
			4 5`,
			`4
			3
			3
			4`,
		},
		{
			`4 2
			1 3 2 6
			3 5
			2 4`,
			`4
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF650D, testCases, 0)
}
