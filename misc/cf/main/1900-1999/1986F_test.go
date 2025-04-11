//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1986/F
// https://codeforces.com/problemset/status/1986/problem/F
func TestCF1986F(t *testing.T) {
	t.Log("Current test is [CF1986F]")
	testCases := [][2]string{
		{
			`6
			2 1
			1 2
			3 3
			1 2
			2 3
			1 3
			5 5
			1 2
			1 3
			3 4
			4 5
			5 3
			6 7
			1 2
			1 3
			2 3
			3 4
			4 5
			4 6
			5 6
			5 5
			1 2
			1 3
			2 3
			2 4
			3 5
			10 12
			1 2
			1 3
			2 3
			2 4
			4 5
			5 6
			6 7
			7 4
			3 8
			8 9
			9 10
			10 8`,
			`0
			3
			4
			6
			6
			21`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1986F, testCases, 0)
}
