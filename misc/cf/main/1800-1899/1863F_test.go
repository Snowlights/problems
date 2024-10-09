//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1863/F
// https://codeforces.com/problemset/status/1863/problem/F
func TestCF1863F(t *testing.T) {
	t.Log("Current test is [CF1863F]")
	testCases := [][2]string{
		{
			`6
			6
			3 2 1 3 7 4
			5
			1 1 1 1 1
			10
			1 2 4 8 4 1 2 3 4 5
			5
			0 0 0 0 0
			5
			1 2 3 0 1
			1
			100500`,
			`111111
			10101
			0001000000
			11111
			11001
			1
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1863F, testCases, 0)
}
