//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1187/E
// https://codeforces.com/problemset/status/1187/problem/E
func TestCF1187E(t *testing.T) {
	t.Log("Current test is [CF1187E]")
	testCases := [][2]string{
		{
			`9
			1 2
			2 3
			2 5
			2 6
			1 4
			4 9
			9 7
			9 8
			`,
			`36`,
		},
		{
			`5
			1 2
			1 3
			2 4
			2 5
			`,
			`14`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1187E, testCases, 0)
}
