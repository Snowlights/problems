//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1775/B
// https://codeforces.com/problemset/status/1775/problem/B
func TestCF1775B(t *testing.T) {
	t.Log("Current test is [CF1775B]")
	testCases := [][2]string{
		{
			`5
			3
			2 1 5
			2 2 4
			2 2 3
			2
			2 1 2
			1 2
			4
			3 1 2 4
			2 2 4
			4 1 2 5 6
			2 2 5
			5
			3 3 1 2
			3 2 5 3
			5 7 2 3 1 4
			5 1 2 6 3 5
			3 2 6 3
			2
			1 1
			1 2`,
			`No
			Yes
			Yes
			Yes
			No
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1775B, testCases, 0)
}
