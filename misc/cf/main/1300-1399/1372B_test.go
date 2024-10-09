//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1372/B
// https://codeforces.com/problemset/status/1372/problem/B
func TestCF1372B(t *testing.T) {
	t.Log("Current test is [CF1372B]")
	testCases := [][2]string{
		{
			`3
			4
			6
			9`,
			`2 2
			3 3
			3 6
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1372B, testCases, 0)
}
