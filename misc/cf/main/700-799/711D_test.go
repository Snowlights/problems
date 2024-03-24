//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/711/D
// https://codeforces.com/problemset/status/711/problem/D
func TestCF711D(t *testing.T) {
	t.Log("Current test is [CF711D]")
	testCases := [][2]string{
		{
			`3
			2 3 1
			`,
			`6`,
		},
		{
			`4
			2 1 1 1
			`,
			`8`,
		},
		{
			`5
			2 4 2 5 3
			`,
			`28`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF711DV2, testCases, 0)
}
