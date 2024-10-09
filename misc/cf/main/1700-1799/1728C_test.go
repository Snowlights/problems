//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1728/C
// https://codeforces.com/problemset/status/1728/problem/C
func TestCF1728C(t *testing.T) {
	t.Log("Current test is [CF1728C]")
	testCases := [][2]string{
		{
			`4
			1
			1
			1000
			4
			1 2 3 4
			3 1 4 2
			3
			2 9 3
			1 100 9
			10
			75019 709259 5 611271314 9024533 81871864 9 3 6 4865
			9503 2 371245467 6 7 37376159 8 364036498 52295554 169`,
			`2
			0
			2
			18
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1728C, testCases, 0)
}
