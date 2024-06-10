//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/911/F
// https://codeforces.com/problemset/status/911/problem/F
func TestCF911F(t *testing.T) {
	t.Log("Current test is [CF911F]")
	testCases := [][2]string{
		{
			`3
			1 2
			1 3
			`,
			`3
			2 3 3
			2 1 1
			`,
		},
		{
			`5
			1 2
			1 3
			2 4
			2 5
			`,
			`9
			3 5 5
			4 3 3
			4 1 1
			4 2 2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF911F, testCases, 0)
}
