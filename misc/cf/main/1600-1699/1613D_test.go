//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1613/D
// https://codeforces.com/problemset/status/1613/problem/D
func TestCF1613D(t *testing.T) {
	t.Log("Current test is [CF1613D]")
	testCases := [][2]string{
		{
			`4
			3
			0 2 1
			2
			1 0
			5
			0 0 0 0 0
			4
			0 1 2 3`,
			`4
			2
			31
			7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1613D, testCases, 0)
}
