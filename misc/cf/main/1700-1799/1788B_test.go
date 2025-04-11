//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1788/B
// https://codeforces.com/problemset/status/1788/problem/B
func TestCF1788B(t *testing.T) {
	t.Log("Current test is [CF1788B]")
	testCases := [][2]string{
		{
			`5
1
161
67
1206
19`,
			`1 0
67 94
60 7
1138 68
14 5`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1788B, testCases, 0)
}
