//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2040/B
// https://codeforces.com/problemset/status/2040/problem/B
func TestCF2040B(t *testing.T) {
	t.Log("Current test is [CF2040B]")
	testCases := [][2]string{
		{
			`4
			1
			2
			4
			20`,
			`1
			2
			2
			4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2040B, testCases, 0)
}
