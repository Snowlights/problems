//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1790/E
// https://codeforces.com/problemset/status/1790/problem/E
func TestCF1790E(t *testing.T) {
	t.Log("Current test is [CF1790E]")
	testCases := [][2]string{
		{
			`6
			2
			5
			10
			6
			18
			36`,
			`3 1
			-1
			13 7
			-1
			25 11
			50 22`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1790E, testCases, 0)
}
