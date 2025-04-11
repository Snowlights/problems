//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1988/D
// https://codeforces.com/problemset/status/1988/problem/D
func TestCF1988D(t *testing.T) {
	t.Log("Current test is [CF1988D]")
	testCases := [][2]string{
		{
			`3
			1
			1000000000000
			5
			47 15 32 29 23
			1 2
			1 3
			2 4
			2 5
			7
			8 10 2 3 5 7 4
			1 2
			1 4
			3 2
			5 3
			6 2
			7 5`,
			`1000000000000
			193
			57`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1988D, testCases, 0)
}
