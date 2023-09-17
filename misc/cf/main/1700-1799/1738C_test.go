//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1738/C
// https://codeforces.com/problemset/status/1738/problem/C
func TestCF1738C(t *testing.T) {
	t.Log("Current test is [CF1738C]")
	testCases := [][2]string{
		{
			`4
			3
			1 3 5
			4
			1 3 5 7
			4
			1 2 3 4
			4
			10 20 30 40`,
			`Alice
			Alice
			Bob
			Alice
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1738C, testCases, 0)
}
