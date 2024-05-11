//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/749/E
// https://codeforces.com/problemset/status/749/problem/E
func TestCF749E(t *testing.T) {
	t.Log("Current test is [CF749E]")
	testCases := [][2]string{
		{
			`3
			2 3 1
			`,
			`1.916666666666666666666666666667`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF749E, testCases, 0)
}
