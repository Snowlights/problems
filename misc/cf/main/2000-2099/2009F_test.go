//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2009/F
// https://codeforces.com/problemset/status/2009/problem/F
func TestCF2009F(t *testing.T) {
	t.Log("Current test is [CF2009F]")
	testCases := [][2]string{
		{
			`5
			3 3
			1 2 3
			1 9
			3 5
			8 8
			5 5
			4 8 3 2 4
			1 14
			3 7
			7 10
			2 11
			1 25
			1 1
			6
			1 1
			5 7
			3 1 6 10 4
			3 21
			6 17
			2 2
			1 5
			1 14
			9 15
			12 13
			5 3
			4 9 10 10 1
			20 25
			3 11
			20 22`,
			`18
			8
			1
			55
			20
			13
			41
			105
			6
			96
			62
			1
			24
			71
			31
			14
			44
			65
			15`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2009F, testCases, 0)
}
