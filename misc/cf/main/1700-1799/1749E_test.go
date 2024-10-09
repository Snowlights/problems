//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/1749/problem/E
// https://codeforces.com/problemset/status/1749/problem/E
func TestCF1749E(t *testing.T) {
	t.Log("Current test is [CF1749E]")
	testCases := [][2]string{
		{
			`4
			2 4
			.#..
			..#.
			3 3
			#.#
			...
			.#.
			5 5
			.....
			.....
			.....
			.....
			.....
			4 3
			#..
			.#.
			#.#
			...`,
			`YES
			.#.#
			#.#.
			NO
			YES
			....#
			...#.
			..#..
			.#...
			#....
			YES
			#..
			.#.
			#.#
			...
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1749E, testCases, 0)
}
