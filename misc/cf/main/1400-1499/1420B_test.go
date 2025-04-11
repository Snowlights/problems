//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1420/B
// https://codeforces.com/problemset/status/1420/problem/B
func TestCF1420B(t *testing.T) {
	t.Log("Current test is [CF1420B]")
	testCases := [][2]string{
		{
			`5
			5
			1 4 3 7 10
			3
			1 1 1
			4
			6 2 5 3
			2
			2 4
			1
			1`,
			`1
			3
			2
			0
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1420B, testCases, 0)
}
