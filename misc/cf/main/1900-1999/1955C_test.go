//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1955/C
// https://codeforces.com/problemset/status/1955/problem/C
func TestCF1955C(t *testing.T) {
	t.Log("Current test is [CF1955C]")
	testCases := [][2]string{
		{
			`6
			4 5
			1 2 4 3
			4 6
			1 2 4 3
			5 20
			2 7 1 8 2
			2 2
			3 2
			2 15
			1 5
			2 7
			5 2`,
			`2
			3
			5
			0
			2
			2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1955C, testCases, 0)
}
