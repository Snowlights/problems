//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1875/D
// https://codeforces.com/problemset/status/1875/problem/D
func TestCF1875D(t *testing.T) {
	t.Log("Current test is [CF1875D]")
	testCases := [][2]string{
		{
			`4
			8
			5 2 1 0 3 0 4 0
			2
			1 2
			5
			1 0 2 114514 0
			8
			0 1 2 0 1 2 0 3`,
			`3
			0
			2
			7
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1875D, testCases, 0)
}
