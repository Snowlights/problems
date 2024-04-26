//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1157/C2
// https://codeforces.com/problemset/status/1157/problem/C2
func TestCF1157C2(t *testing.T) {
	t.Log("Current test is [CF1157C2]")
	testCases := [][2]string{
		{
			`5
			1 2 4 3 2`,
			`4
			LRRR`,
		},
		{
			`7
			1 3 5 6 5 4 2`,
			`6
			LRLRRR`,
		},
		{
			`3
			2 2 2`,
			`1
			R`,
		},
		{
			`4
			1 2 4 3`,
			`4
			LLRR`,
		},
	}

	codeforces.AssertEqualStringCase(t, CF1157C2, testCases, 0)
}
