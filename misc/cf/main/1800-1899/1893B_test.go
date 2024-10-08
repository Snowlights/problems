//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1893/B
// https://codeforces.com/problemset/status/1893/problem/B
func TestCF1893B(t *testing.T) {
	t.Log("Current test is [CF1893B]")
	testCases := [][2]string{
		{
			`7
			2 1
			6 4
			5
			5 5
			1 7 2 4 5
			5 4 1 2 7
			1 9
			7
			1 2 3 4 5 6 7 8 9
			3 2
			1 3 5
			2 4
			10 5
			1 9 2 3 8 1 4 7 2 9
			7 8 5 4 6
			2 1
			2 2
			1
			6 1
			1 1 1 1 1 1
			777`,
			`6 5 4
			1 1 7 7 2 2 4 4 5 5
			9 8 7 7 6 5 4 3 2 1
			1 3 5 2 4
			1 9 2 3 8 8 1 4 4 7 7 2 9 6 5
			2 2 1
			777 1 1 1 1 1 1
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1893B, testCases, 0)
}
