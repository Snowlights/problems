//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2023/B
// https://codeforces.com/problemset/status/2023/problem/B
func TestCF2023B(t *testing.T) {
	t.Log("Current test is [CF2023B]")
	testCases := [][2]string{
		{
			`4
			2
			15 16
			2 1
			5
			10 10 100 100 1000
			3 4 1 1 1
			3
			100 49 50
			3 2 2
			4
			100 200 300 1000
			2 3 4 1`,
			`16
			200
			100
			1000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2023B, testCases, 0)
}
