//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1443/B
// https://codeforces.com/problemset/status/1443/problem/B
func TestCF1443B(t *testing.T) {
	t.Log("Current test is [CF1443B]")
	testCases := [][2]string{
		{
			`2
			1 1
			01000010
			5 1
			01101110
			`,
			`2
			6
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1443B, testCases, 0)
}
