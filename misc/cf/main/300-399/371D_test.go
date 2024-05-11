//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/371/D
// https://codeforces.com/problemset/status/371/problem/D
func TestCF371D(t *testing.T) {
	t.Log("Current test is [CF371D]")
	testCases := [][2]string{
		{
			`2
			5 10
			6
			1 1 4
			2 1
			1 2 5
			1 1 4
			2 1
			2 2
			`,
			`4
			5
			8
			`,
		},
		{
			`3
			5 10 8
			6
			1 1 12
			2 2
			1 1 6
			1 3 2
			2 2
			2 3
			`,
			`7
			10
			5
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF371D, testCases, 0)
}
