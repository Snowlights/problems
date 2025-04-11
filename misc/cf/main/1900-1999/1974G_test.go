//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1974/G
// https://codeforces.com/problemset/status/1974/problem/G
func TestCF1974G(t *testing.T) {
	t.Log("Current test is [CF1974G]")
	testCases := [][2]string{
		{
			`6
			3 3
			2 2 2
			6 5
			2 2 8 2 6 8
			6 4
			4 10 3 8 6 10
			2 1
			1 1
			4 1
			4 1 3 1
			4 2
			1 3 4 3`,
			`2
			4
			3
			1
			2
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1974G, testCases, 0)
}
