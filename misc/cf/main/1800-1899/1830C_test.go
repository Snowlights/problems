//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1830/C
// https://codeforces.com/problemset/status/1830/problem/C
func TestCF1830C(t *testing.T) {
	t.Log("Current test is [CF1830C]")
	testCases := [][2]string{
		{
			`7
			6 0
			5 0
			8 1
			1 3
			10 2
			3 4
			6 9
			1000 3
			100 701
			200 801
			300 901
			28 5
			1 12
			3 20
			11 14
			4 9
			18 19
			4 3
			1 4
			1 4
			1 4`,
			`5
			0
			0
			4
			839415253
			140
			2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1830C, testCases, 0)
}
