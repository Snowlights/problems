//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1744/F
// https://codeforces.com/problemset/status/1744/problem/F
func TestCF1744F(t *testing.T) {
	t.Log("Current test is [CF1744F]")
	testCases := [][2]string{
		{
			`8
			1
			0
			2
			1 0
			3
			1 0 2
			4
			0 2 1 3
			5
			3 1 0 2 4
			6
			2 0 4 1 3 5
			8
			3 7 2 6 0 1 5 4
			4
			2 0 1 3`,
			`1
			2
			4
			4
			8
			8
			15
			6
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1744F, testCases, 0)
}
