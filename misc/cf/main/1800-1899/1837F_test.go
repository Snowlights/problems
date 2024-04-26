//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1837/F
// https://codeforces.com/problemset/status/1837/problem/F
func TestCF1837F(t *testing.T) {
	t.Log("Current test is [CF1837F]")
	testCases := [][2]string{
		{
			`6
			5 4
			1 10 1 1 1
			5 3
			1 20 5 15 3
			5 3
			1 20 3 15 5
			10 6
			10 8 20 14 3 8 6 4 16 11
			10 5
			9 9 2 13 15 19 4 9 13 12
			1 1
			1`,
			`2
			6
			5
			21
			18
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1837F, testCases, 0)
}
