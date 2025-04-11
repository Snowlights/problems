//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2074/G
// https://codeforces.com/problemset/status/2074/problem/G
func TestCF2074G(t *testing.T) {
	t.Log("Current test is [CF2074G]")
	testCases := [][2]string{
		{
			`6
			3
			1 2 3
			4
			2 1 3 4
			6
			2 1 2 1 1 1
			6
			1 2 1 3 1 5
			9
			9 9 8 2 4 4 3 5 3
			9
			9 9 3 2 4 4 8 5 3`,
			`6
			24
			5
			30
			732
			696`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2074G, testCases, 0)
}
