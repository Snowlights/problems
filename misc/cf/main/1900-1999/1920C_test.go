//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1920/C
// https://codeforces.com/problemset/status/1920/problem/C
func TestCF1920C(t *testing.T) {
	t.Log("Current test is [CF1920C]")
	testCases := [][2]string{
		{
			`8
			4
			1 2 1 4
			3
			1 2 3
			5
			1 1 1 1 1
			6
			1 3 1 1 3 1
			6
			6 2 6 2 2 2
			6
			2 6 3 6 6 6
			10
			1 7 5 1 4 3 1 3 1 4
			1
			1`,
			`2
			1
			2
			4
			4
			1
			2
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1920C, testCases, 0)
}
