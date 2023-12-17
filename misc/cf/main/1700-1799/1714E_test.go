//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1714/E
// https://codeforces.com/problemset/status/1714/problem/E
func TestCF1714E(t *testing.T) {
	t.Log("Current test is [CF1714E]")
	testCases := [][2]string{
		{
			`10
			2
			6 11
			3
			2 18 22
			5
			5 10 5 10 5
			4
			1 2 4 8
			2
			4 5
			3
			93 96 102
			2
			40 6
			2
			50 30
			2
			22 44
			2
			1 5`,
			`Yes
			No
			Yes
			Yes
			No
			Yes
			No
			No
			Yes
			No
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1714E, testCases, 0)
}
