//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/193/A
// https://codeforces.com/problemset/status/193/problem/A
func TestCF193A(t *testing.T) {
	t.Log("Current test is [CF193A]")
	testCases := [][2]string{
		{
			`5 4
			####
			#..#
			#..#
			#..#
			####`,
			`2`,
		},
		{
			`5 5
			#####
			#...#
			#####
			#...#
			#####`,
			`2`,
		},
		{
			`3 3
			.#.
			###
			.#.`,
			`1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF193A, testCases, 0)
}
