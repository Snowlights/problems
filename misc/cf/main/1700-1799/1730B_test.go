//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1730/B
// https://codeforces.com/problemset/status/1730/problem/B
func TestCF1730B(t *testing.T) {
	t.Log("Current test is [CF1730B]")
	testCases := [][2]string{
		{
			`7
			1
			0
			3
			2
			3 1
			0 0
			2
			1 4
			0 0
			3
			1 2 3
			0 0 0
			3
			1 2 3
			4 1 2
			3
			3 3 3
			5 3 3
			6
			5 4 7 2 10 4
			3 2 5 1 4 6`,
			`0
			2
			2.5
			2
			1
			3
			6
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1730B, testCases, 0)
}
