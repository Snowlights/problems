//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1730/E
// https://codeforces.com/problemset/status/1730/problem/E
func TestCF1730E(t *testing.T) {
	t.Log("Current test is [CF1730E]")
	testCases := [][2]string{
		{
			`6
			1
			1
			2
			2 4
			2
			2 3
			4
			2 4 7 14
			7
			16 5 18 7 7 12 14
			6
			16 14 2 6 16 2`,
			`1
			3
			2
			7
			10
			19
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1730E, testCases, 0)
}
