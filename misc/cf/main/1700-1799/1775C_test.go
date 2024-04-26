//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1775/C
// https://codeforces.com/problemset/status/1775/problem/C
func TestCF1775C(t *testing.T) {
	t.Log("Current test is [CF1775C]")
	testCases := [][2]string{
		{
			`5
			10 8
			10 10
			10 42
			20 16
			1000000000000000000 0`,
			`12
			10
			-1
			24
			1152921504606846976
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1775C, testCases, 0)
}
