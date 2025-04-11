//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2046/A
// https://codeforces.com/problemset/status/2046/problem/A
func TestCF2046A(t *testing.T) {
	t.Log("Current test is [CF2046A]")
	testCases := [][2]string{
		{
			`3
			1
			-10
			5
			3
			1 2 3
			10 -5 -3
			4
			2 8 5 3
			1 10 3 4`,
			`-5
			16
			29`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2046A, testCases, 0)
}
