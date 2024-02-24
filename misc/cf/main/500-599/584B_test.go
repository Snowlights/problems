//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/584/B
// https://codeforces.com/problemset/status/584/problem/B
func TestCF584B(t *testing.T) {
	t.Log("Current test is [CF584B]")
	testCases := [][2]string{
		{
			`1`,
			`20`,
		},
		{
			`2`,
			`680`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF584B, testCases, 0)
}
