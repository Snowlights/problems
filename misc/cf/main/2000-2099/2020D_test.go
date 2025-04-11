//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2020/D
// https://codeforces.com/problemset/status/2020/problem/D
func TestCF2020D(t *testing.T) {
	t.Log("Current test is [CF2020D]")
	testCases := [][2]string{
		{
			`3
			10 2
			1 2 4
			2 2 4
			100 1
			19 2 4
			100 3
			1 2 5
			7 2 6
			17 2 31`,
			`2
			96
			61`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2020D, testCases, 0)
}
