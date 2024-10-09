//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/463/D
// https://codeforces.com/problemset/status/463/problem/D
func TestCF463D(t *testing.T) {
	t.Log("Current test is [CF463D]")
	testCases := [][2]string{
		{
			`4 3
			1 4 2 3
			4 1 2 3
			1 2 4 3
			`,
			`3`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF463D, testCases, 0)
}
