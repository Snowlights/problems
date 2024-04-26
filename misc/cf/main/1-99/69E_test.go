//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/69/E
// https://codeforces.com/problemset/status/69/problem/E
func TestCF69E(t *testing.T) {
	t.Log("Current test is [CF69E]")
	testCases := [][2]string{
		{
			`5 3
			1
			2
			2
			3
			3
			`,
			`1
			3
			2
			`,
		},
		{
			`6 4
			3
			3
			3
			4
			4
			2
			`,
			`4
			Nothing
			3
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF69E, testCases, 0)
}
