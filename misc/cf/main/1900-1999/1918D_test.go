//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1918/D
// https://codeforces.com/problemset/status/1918/problem/D
func TestCF1918D(t *testing.T) {
	t.Log("Current test is [CF1918D]")
	testCases := [][2]string{
		{
			`3
			6
			1 4 5 3 3 2
			5
			1 2 3 4 5
			6
			4 1 6 3 10 7`,
			`7
			5
			11`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1918D, testCases, 0)
}
