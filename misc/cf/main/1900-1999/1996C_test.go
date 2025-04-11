//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1996/C
// https://codeforces.com/problemset/status/1996/problem/C
func TestCF1996C(t *testing.T) {
	t.Log("Current test is [CF1996C]")
	testCases := [][2]string{
		{
			`3
			5 3
			abcde
			edcba
			1 5
			1 4
			3 3
			4 2
			zzde
			azbe
			1 3
			1 4
			6 3
			uwuwuw
			wuwuwu
			2 4
			1 3
			1 6`,
			`0
			1
			0
			2
			2
			1
			1
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1996C, testCases, 0)
}
