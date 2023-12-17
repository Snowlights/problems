//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1443/C
// https://codeforces.com/problemset/status/1443/problem/C
func TestCF1443C(t *testing.T) {
	t.Log("Current test is [CF1443C]")
	testCases := [][2]string{
		{
			`4
			4
			3 7 4 5
			2 1 2 4
			4
			1 2 3 4
			3 3 3 3
			2
			1 2
			10 10
			2
			10 10
			1 2`,
			`5
			3
			2
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1443C, testCases, 0)
}
