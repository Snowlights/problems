//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1169/B
// https://codeforces.com/problemset/status/1169/problem/B
func TestCF1169B(t *testing.T) {
	t.Log("Current test is [CF1169B]")
	testCases := [][2]string{
		{
			`4 6
			1 2
			1 3
			1 4
			2 3
			2 4
			3 4
			`,
			`NO`,
		},
		{
			`5 4
			1 2
			2 3
			3 4
			4 5
			`,
			`YES`,
		},
		{
			`300000 5
			1 2
			1 2
			1 2
			1 2
			1 2
			`,
			`YES`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1169B, testCases, 0)
}
