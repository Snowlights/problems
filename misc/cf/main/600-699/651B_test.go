//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/651/B
// https://codeforces.com/problemset/status/651/problem/B
func TestCF651B(t *testing.T) {
	t.Log("Current test is [CF651B]")
	testCases := [][2]string{
		{
			`5
			20 30 10 50 40`,
			`4`,
		},
		{
			`4
			200 100 100 200`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF651B, testCases, 0)
}
