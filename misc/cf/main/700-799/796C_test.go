//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/796/C
// https://codeforces.com/problemset/status/796/problem/C
func TestCF796C(t *testing.T) {
	t.Log("Current test is [CF796C]")
	testCases := [][2]string{
		{
			`5
			1 2 3 4 5
			1 2
			2 3
			3 4
			4 5`,
			`5`,
		},
		{
			`7
			38 -29 87 93 39 28 -55
			1 2
			2 5
			3 2
			2 4
			1 7
			7 6`,
			`93`,
		},
		{
			`5
			1 2 7 6 7
			1 5
			5 3
			3 4
			2 4`,
			`8`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF796C, testCases, 0)
}
