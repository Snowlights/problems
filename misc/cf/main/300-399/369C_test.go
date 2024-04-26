//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/369/C
// https://codeforces.com/problemset/status/369/problem/C
func TestCF369C(t *testing.T) {
	t.Log("Current test is [CF369C]")
	testCases := [][2]string{
		{
			`5
			1 2 2
			2 3 2
			3 4 2
			4 5 2
			`,
			`1
			5`,
		},
		{
			`5
			1 2 1
			2 3 2
			2 4 1
			4 5 1
			`,
			`1
			3 
			`,
		},
		{
			`5
			1 2 2
			1 3 2
			1 4 2
			1 5 2
			`,
			`4
			5 4 3 2 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF369C, testCases, 0)
}
