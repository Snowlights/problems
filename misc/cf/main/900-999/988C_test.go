//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/988/C
// https://codeforces.com/problemset/status/988/problem/C
func TestCF988C(t *testing.T) {
	t.Log("Current test is [CF988C]")
	testCases := [][2]string{
		{
			`2
			5
			2 3 1 3 2
			6
			1 1 2 2 2 1
			`,
			`YES
			2 6
			1 2
			`,
		},
		{
			`3
			1
			5
			5
			1 1 1 1 1
			2
			2 3
			`,
			`NO`,
		},
		{
			`4
			6
			2 2 2 2 2 2
			5
			2 2 2 2 2
			3
			2 2 2
			5
			2 2 2 2 2
			`,
			`YES
			2 2
			4 1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF988C, testCases, 0)
}
