//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1031/B
// https://codeforces.com/problemset/status/1031/problem/B
func TestCF1031B(t *testing.T) {
	t.Log("Current test is [CF1031B]")
	testCases := [][2]string{
		{
			`4
			3 3 2
			1 2 0`,
			`YES
			1 3 2 0`,
		},
		{
			`3
			1 3
			3 2`,
			`NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1031B, testCases, 0)
}
