//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1065/C
// https://codeforces.com/problemset/status/1065/problem/C
func TestCF1065C(t *testing.T) {
	t.Log("Current test is [CF1065C]")
	testCases := [][2]string{
		{
			`5 5
			3 1 2 2 4
			`,
			`2
`,
		},
		{
			`4 5
			2 3 4 5
			`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1065C, testCases, 0)
}
