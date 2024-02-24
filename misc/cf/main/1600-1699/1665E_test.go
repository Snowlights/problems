//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1665/E
// https://codeforces.com/problemset/status/1665/problem/E
func TestCF1665E(t *testing.T) {
	t.Log("Current test is [CF1665E]")
	testCases := [][2]string{
		{
			`2
			5
			6 1 3 2 1
			4
			1 2
			2 3
			2 4
			2 5
			4
			0 2 1 1073741823
			4
			1 2
			2 3
			1 3
			3 4`,
			`7
			3
			3
			1
			2
			3
			1
			1073741823
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1665E, testCases, 0)
}
