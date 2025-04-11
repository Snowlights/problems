//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/937/B
// https://codeforces.com/problemset/status/937/problem/B
func TestCF937B(t *testing.T) {
	t.Log("Current test is [CF937B]")
	testCases := [][2]string{
		{
			`3 6`,
			`5`,
		},
		{
			`3 4`,
			`-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF937B, testCases, 0)
}
