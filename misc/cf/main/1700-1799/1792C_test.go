//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1792/C
// https://codeforces.com/problemset/status/1792/problem/C
func TestCF1792C(t *testing.T) {
	t.Log("Current test is [CF1792C]")
	testCases := [][2]string{
		{
			`4
			5
			1 5 4 2 3
			3
			1 2 3
			4
			2 1 4 3
			6
			5 2 4 1 6 3`,
			`2
			0
			1
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1792C, testCases, 0)
}
