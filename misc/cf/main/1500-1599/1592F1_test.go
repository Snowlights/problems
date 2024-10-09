//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1592/F1
// https://codeforces.com/problemset/status/1592/problem/F1
func TestCF1592F1(t *testing.T) {
	t.Log("Current test is [CF1592F1]")
	testCases := [][2]string{
		{
			`3 3
			WWW
			WBB
			WBB`,
			`3`,
		},
		{
			`10 15
			WWWBBBWBBBBBWWW
			BBBBWWWBBWWWBBB
			BBBWWBWBBBWWWBB
			BBWBWBBWWWBBWBW
			BBBBWWWBBBWWWBB
			BWBBWWBBBBBBWWW
			WBWWBBBBWWBBBWW
			WWBWWWWBBWWBWWW
			BWBWWBWWWWWWBWB
			BBBWBWBWBBBWWBW`,
			`74`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1592F1, testCases, 0)
}
