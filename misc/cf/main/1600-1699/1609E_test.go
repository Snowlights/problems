//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1609/E
// https://codeforces.com/problemset/status/1609/problem/E
func TestCF1609E(t *testing.T) {
	t.Log("Current test is [CF1609E]")
	testCases := [][2]string{
		{
			`9 12
			aaabccccc
			4 a
			4 b
			2 b
			5 a
			1 b
			6 b
			5 c
			2 a
			1 a
			5 a
			6 b
			7 b`,
			`0
			1
			2
			2
			1
			2
			1
			2
			2
			2
			2
			2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1609E, testCases, 0)
}
