//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1354/B
// https://codeforces.com/problemset/status/1354/problem/B
func TestCF1354B(t *testing.T) {
	t.Log("Current test is [CF1354B]")
	testCases := [][2]string{
		{
			`7
			123
			12222133333332
			112233
			332211
			12121212
			333333
			31121`,
			`3
			3
			4
			4
			0
			0
			4
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1354B, testCases, 0)
}
