//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1900/C
// https://codeforces.com/problemset/status/1900/problem/C
func TestCF1900C(t *testing.T) {
	t.Log("Current test is [CF1900C]")
	testCases := [][2]string{
		{
			`5
			3
			LRU
			2 3
			0 0
			0 0
			3
			ULR
			3 2
			0 0
			0 0
			2
			LU
			0 2
			0 0
			4
			RULR
			3 0
			0 0
			0 4
			2 0
			7
			LLRRRLU
			5 2
			3 6
			0 0
			7 0
			4 0
			0 0
			0 0`,
			`0
			1
			1
			3
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1900C, testCases, 0)
}
