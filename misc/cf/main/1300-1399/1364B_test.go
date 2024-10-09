//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1364/B
// https://codeforces.com/problemset/status/1364/problem/B
func TestCF1364B(t *testing.T) {
	t.Log("Current test is [CF1364B]")
	testCases := [][2]string{
		{
			`2
			3
			3 2 1
			4
			1 3 4 2
			`,
			`2
			3 1 
			3
			1 4 2 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1364B, testCases, 0)
}
