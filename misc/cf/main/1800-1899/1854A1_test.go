//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1854/A1
// https://codeforces.com/problemset/status/1854/problem/A1
func TestCF1854A1(t *testing.T) {
	t.Log("Current test is [CF1854A1]")
	testCases := [][2]string{
		{
			`10
			2
			2 1
			4
			1 2 -10 3
			5
			2 1 1 1 1
			8
			0 0 0 0 0 0 0 0
			5
			1 2 -4 3 -10
			10
			11 12 13 14 15 -15 -16 -17 -18 -19
			7
			1 9 3 -4 -3 -2 -1
			3
			10 9 8
			20
			1 -14 2 -10 6 -5 10 -13 10 7 -14 19 -5 19 1 18 -16 -7 12 8
			20
			-15 -17 -13 8 14 -13 10 -4 11 -4 -16 -6 15 -4 -2 7 -9 5 -5 17`,
			`1
			2 1
			3
			4 4
			4 4
			3 4
			4
			2 1
			3 1
			4 1
			5 1
			0
			7
			3 4
			3 4
			5 4
			5 4
			5 4
			5 4
			5 4
			15
			6 1
			6 1
			6 1
			7 2
			7 2
			7 2
			8 3
			8 3
			8 3
			9 4
			9 4
			9 4
			10 5
			10 5
			10 5
			8
			3 4
			3 4
			2 4
			2 4
			2 4
			2 4
			1 4
			1 4
			3
			2 1
			3 1
			3 1
			31
			14 1
			18 7
			13 11
			15 11
			6 4
			5 17
			19 6
			19 12
			10 5
			11 12
			1 17
			15 19
			16 10
			14 2
			16 11
			20 7
			7 6
			9 5
			3 6
			6 14
			17 18
			18 14
			12 3
			17 16
			8 18
			13 16
			9 8
			14 8
			16 2
			11 8
			12 7
			31
			5 12
			19 13
			9 1
			5 17
			18 19
			6 16
			15 8
			6 9
			15 14
			7 10
			19 7
			17 20
			14 4
			15 20
			4 3
			1 8
			16 12
			16 15
			5 6
			12 10
			11 15
			20 3
			20 19
			13 14
			11 14
			18 10
			7 3
			12 17
			4 7
			13 2
			11 13
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1854A1, testCases, 0)
}
