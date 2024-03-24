//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1747/C
// https://codeforces.com/problemset/status/1747/problem/C
func TestCF1747C(t *testing.T) {
	t.Log("Current test is [CF1747C]")
	testCases := [][2]string{
		{
			`3
			2
			1 1
			2
			2 1
			3
			5 4 4`,
			`Bob
			Alice
			Alice
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1747C, testCases, 0)
}
