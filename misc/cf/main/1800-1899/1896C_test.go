//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1896/C
// https://codeforces.com/problemset/status/1896/problem/C
func TestCF1896C(t *testing.T) {
	t.Log("Current test is [CF1896C]")
	testCases := [][2]string{
		{
			`7
			1 0
			1
			2
			1 1
			1
			2
			3 0
			2 4 3
			4 1 2
			3 1
			2 4 3
			4 1 2
			3 2
			2 4 3
			4 1 2
			3 3
			2 4 3
			4 1 2
			5 2
			6 4 5 6 2
			9 7 9 1 1`,
			`YES
			2
			NO
			NO
			YES
			2 4 1
			YES
			4 1 2
			NO
			YES
			1 9 9 7 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1896C, testCases, 0)
}
