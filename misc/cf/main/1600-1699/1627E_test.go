//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1627/E
// https://codeforces.com/problemset/status/1627/problem/E
func TestCF1627E(t *testing.T) {
	t.Log("Current test is [CF1627E]")
	testCases := [][2]string{
		{
			`4
			5 3 3
			5 17 8 1 4
			1 3 3 3 4
			3 1 5 2 5
			3 2 5 1 6
			6 3 3
			5 17 8 1 4 2
			1 3 3 3 4
			3 1 5 2 5
			3 2 5 1 6
			5 3 1
			5 17 8 1 4
			1 3 5 3 100
			5 5 5
			3 2 3 7 5
			3 5 4 2 1
			2 2 5 4 5
			4 4 5 2 3
			1 2 4 2 2
			3 3 5 2 4`,
			`16
			NO ESCAPE
			-90
			27`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1627E, testCases, 0)
}
