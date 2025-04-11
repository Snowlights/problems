//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/817/E
// https://codeforces.com/problemset/status/817/problem/E
func TestCF817E(t *testing.T) {
	t.Log("Current test is [CF817E]")
	testCases := [][2]string{
		{
			`5
			1 3
			1 4
			3 6 3
			2 4
			3 6 3`,
			`1
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF817E, testCases, 0)
}
