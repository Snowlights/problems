//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1821/E
// https://codeforces.com/problemset/status/1821/problem/E
func TestCF1821E(t *testing.T) {
	t.Log("Current test is [CF1821E]")
	testCases := [][2]string{
		{
			`7
			0
			()
			0
			(())
			1
			(())
			5
			()
			1
			(()()(()))
			2
			((())()(()())((())))
			3
			((())()(()())((())))`,
			`0
			1
			0
			0
			1
			4
			2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1821E, testCases, 0)
}
