//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1759/D
// https://codeforces.com/problemset/status/1759/problem/D
func TestCF1759D(t *testing.T) {
	t.Log("Current test is [CF1759D]")
	testCases := [][2]string{
		{
			`10
			6 11
			5 43
			13 5
			4 16
			10050 12345
			2 6
			4 30
			25 10
			2 81
			1 7`,
			`60
			200
			65
			60
			120600000
			10
			100
			200
			100
			7
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1759D, testCases, 0)
}
