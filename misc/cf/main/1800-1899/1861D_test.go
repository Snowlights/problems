//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1861/problem/D
// https://codeforces.com/problemset/status/1861/problem/D
func TestCF1861D(t *testing.T) {
	t.Log("Current test is [CF1861D]")
	testCases := [][2]string{
		{
			`3
			5
			1 1 2 2 2
			6
			5 4 3 2 5 1
			3
			1 2 3`,
			`3
			2
			0
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1861D, testCases, 0)
}
