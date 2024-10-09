//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/455/C
// https://codeforces.com/problemset/status/455/problem/C
func TestCF455C(t *testing.T) {
	t.Log("Current test is [CF455C]")
	testCases := [][2]string{
		{
			`6 0 6
			2 1 2
			2 3 4
			2 5 6
			2 3 2
			2 5 3
			1 1`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF455C, testCases, 0)
}
