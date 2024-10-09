//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1872/G
// https://codeforces.com/problemset/status/1872/problem/G
func TestCF1872G(t *testing.T) {
	t.Log("Current test is [CF1872G]")
	testCases := [][2]string{
		{
			`9
			4
			1 3 1 3
			4
			1 1 2 3
			5
			1 1 1 1 1
			5
			10 1 10 1 10
			1
			1
			2
			2 2
			3
			2 1 2
			4
			2 1 1 3
			6
			2 1 2 1 1 3`,
			`2 4
			3 4
			1 1
			1 5
			1 1
			1 2
			2 2
			4 4
			1 6
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1872G, testCases, 0)
}
