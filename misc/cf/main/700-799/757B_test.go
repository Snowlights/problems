//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/757/B
// https://codeforces.com/problemset/status/757/problem/B
func TestCF757B(t *testing.T) {
	t.Log("Current test is [CF757B]")
	testCases := [][2]string{
		{
			`3
			2 3 4`,
			`2`,
		},
		{
			`5
			2 3 4 6 7`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF757B, testCases, 0)
}
