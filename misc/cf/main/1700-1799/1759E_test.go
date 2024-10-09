//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1759/E
// https://codeforces.com/problemset/status/1759/problem/E
func TestCF1759E(t *testing.T) {
	t.Log("Current test is [CF1759E]")
	testCases := [][2]string{
		{
			`8
			4 1
			2 1 8 9
			3 3
			6 2 60
			4 5
			5 1 100 5
			3 2
			38 6 3
			1 1
			12
			4 6
			12 12 36 100
			4 1
			2 1 1 15
			3 5
			15 1 13`,
			`4
			3
			3
			3
			0
			4
			4
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1759E, testCases, 0)
}
