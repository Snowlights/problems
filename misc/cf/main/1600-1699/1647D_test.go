//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1647/D
// https://codeforces.com/problemset/status/1647/problem/D
func TestCF1647D(t *testing.T) {
	t.Log("Current test is [CF1647D]")
	testCases := [][2]string{
		{
			`8
			6 2
			12 2
			36 2
			8 2
			1000 10
			2376 6
			128 4
			16384 4`,
			`NO
			NO
			YES
			NO
			YES
			YES
			NO
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1647D, testCases, 0)
}
