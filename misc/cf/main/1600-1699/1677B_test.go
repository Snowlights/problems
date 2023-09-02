//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1677/B
// https://codeforces.com/problemset/status/1677/problem/B
func TestCF1677B(t *testing.T) {
	t.Log("Current test is [CF1677B]")
	testCases := [][2]string{
		{
			`3
			2 2
			1100
			4 2
			11001101
			2 4
			11001101`,
			`2 3 4 3
			2 3 4 3 5 4 6 5
			2 3 3 3 4 4 4 5
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1677B, testCases, 0)
}
