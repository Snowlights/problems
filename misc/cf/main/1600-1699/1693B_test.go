//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1693/B
// https://codeforces.com/problemset/status/1693/problem/B
func TestCF1693B(t *testing.T) {
	t.Log("Current test is [CF1693B]")
	testCases := [][2]string{
		{
			`4
			2
			1
			1 5
			2 9
			3
			1 1
			4 5
			2 4
			6 10
			4
			1 2 1
			6 9
			5 6
			4 5
			2 4
			5
			1 2 3 4
			5 5
			4 4
			3 3
			2 2
			1 1`,
			`1
			2
			2
			5
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1693B, testCases, 0)
}
