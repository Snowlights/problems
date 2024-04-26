//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1198/B
// https://codeforces.com/problemset/status/1198/problem/B
func TestCF1198B(t *testing.T) {
	t.Log("Current test is [CF1198B]")
	testCases := [][2]string{
		{
			`4
			1 2 3 4
			3
			2 3
			1 2 2
			2 1
			`,
			`3 2 3 4`,
		},
		{
			`5
			3 50 2 1 10
			3
			1 2 0
			2 8
			1 3 20
			`,
			`8 8 20 8 10`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1198B, testCases, 0)
}
