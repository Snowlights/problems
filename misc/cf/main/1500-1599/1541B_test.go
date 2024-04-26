//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1541/B
// https://codeforces.com/problemset/status/1541/problem/B
func TestCF1541B(t *testing.T) {
	t.Log("Current test is [CF1541B]")
	testCases := [][2]string{
		{
			`3
			2
			3 1
			3
			6 1 5
			5
			3 1 5 9 2
			`,
			`1
			1
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1541B, testCases, 0)
}
