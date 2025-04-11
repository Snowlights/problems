//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/70/D
// https://codeforces.com/problemset/status/70/problem/D
func TestCF70D(t *testing.T) {
	t.Log("Current test is [CF70D]")
	testCases := [][2]string{
		{
			`8
			1 0 0
			1 2 0
			1 2 2
			2 1 0
			1 0 2
			2 1 1
			2 2 1
			2 20 -1`,
			`YES
			YES
			YES
			NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF70D, testCases, 0)
}
