//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2060/E
// https://codeforces.com/problemset/status/2060/problem/E
func TestCF2060E(t *testing.T) {
	t.Log("Current test is [CF2060E]")
	testCases := [][2]string{
		{
			`5
			3 2 1
			1 2
			2 3
			1 3
			2 1 1
			1 2
			1 2
			3 2 0
			3 2
			1 2
			1 0 0
			3 3 1
			1 2
			1 3
			2 3
			1 2`,
			`3
			0
			2
			0
			2`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF2060E, testCases, 0)
}
