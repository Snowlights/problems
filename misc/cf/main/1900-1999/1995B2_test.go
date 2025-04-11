//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1995/B2
// https://codeforces.com/problemset/status/1995/problem/B2
func TestCF1995B2(t *testing.T) {
	t.Log("Current test is [CF1995B2]")
	testCases := [][2]string{
		{
			`7
			3 10
			1 2 3
			2 2 1
			3 1033
			206 207 1000
			3 4 1
			6 20
			4 2 7 5 6 1
			1 2 1 3 1 7
			8 100000
			239 30 610 122 24 40 8 2
			12 13123 112 1456 124 100 123 10982
			6 13
			2 4 11 1 3 5
			2 2 1 2 2 1
			8 10330
			206 210 200 201 198 199 222 1000
			9 10 11 12 13 14 15 16
			2 10000000000
			11 12
			87312315 753297050`,
			`7
			1033
			19
			99990
			13
			10000
			9999999999`,
		},
		{
			`1
			6 20
			4 2 7 5 6 1
			1 2 1 3 1 7`,
			`19`,
		},
		{
			`1
			2 2
			3 4
			2 5`,
			`0`,
		},
		{
			`1
			2 5
			3 4
			3 3`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1995B2, testCases, 0)
}
