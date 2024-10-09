//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1771/B
// https://codeforces.com/problemset/status/1771/problem/B
func TestCF1771B(t *testing.T) {
	t.Log("Current test is [CF1771B]")
	testCases := [][2]string{
		{
			`2
			3 2
			1 3
			2 3
			4 2
			1 2
			2 3`,
			`4
			5
			`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF1771B, testCases, 0)
}
