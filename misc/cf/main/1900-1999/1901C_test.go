//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1901/C
// https://codeforces.com/problemset/status/1901/problem/C
func TestCF1901C(t *testing.T) {
	t.Log("Current test is [CF1901C]")
	testCases := [][2]string{
		{
			`4
			1
			10
			2
			4 6
			6
			2 1 2 1 2 1
			2
			0 32`,
			`0
			2
			2 5
			1
			1
			6
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1901C, testCases, 0)
}
