//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1610/C
// https://codeforces.com/problemset/status/1610/problem/C
func TestCF1610C(t *testing.T) {
	t.Log("Current test is [CF1610C]")
	testCases := [][2]string{
		{
			`3
			3
			1 2
			2 1
			1 1
			2
			0 0
			0 1
			2
			1 0
			0 1
			`,
			`2
			1
			2
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1610C, testCases, 0)
}
