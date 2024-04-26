//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/845/C
// https://codeforces.com/problemset/status/845/problem/C
func TestCF845C(t *testing.T) {
	t.Log("Current test is [CF845C]")
	testCases := [][2]string{
		{
			`3
			1 2
			2 3
			4 5
			`,
			`YES`,
		},
		{
			`4
			1 2
			2 3
			2 3
			1 2
			`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF845C, testCases, 0)
}
