//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1738/B
// https://codeforces.com/problemset/status/1738/problem/B
func TestCF1738B(t *testing.T) {
	t.Log("Current test is [CF1738B]")
	testCases := [][2]string{
		{
			`4
			5 5
			1 2 3 4 5
			7 4
			-6 -5 -3 0
			3 3
			2 3 4
			3 2
			3 4`,
			`Yes
			Yes
			No
			No
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1738B, testCases, 0)
}
