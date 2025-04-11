//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2021/B
// https://codeforces.com/problemset/status/2021/problem/B
func TestCF2021B(t *testing.T) {
	t.Log("Current test is [CF2021B]")
	testCases := [][2]string{
		{
			`3
			6 3
			0 3 2 1 5 2
			6 2
			1 3 4 1 0 2
			4 5
			2 5 10 3`,
			`4
			6
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2021B, testCases, 0)
}
