//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2065/D
// https://codeforces.com/problemset/status/2065/problem/D
func TestCF2065D(t *testing.T) {
	t.Log("Current test is [CF2065D]")
	testCases := [][2]string{
		{
			`3
			2 2
			4 4
			6 1
			3 4
			2 2 2 2
			3 2 1 2
			4 1 2 1
			2 3
			3 4 5
			1 1 9`,
			`41
			162
			72`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2065D, testCases, 0)
}
