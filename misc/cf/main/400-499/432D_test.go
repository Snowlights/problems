//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/432/D
// https://codeforces.com/problemset/status/432/problem/D
func TestCF432D(t *testing.T) {
	t.Log("Current test is [CF432D]")
	testCases := [][2]string{
		{
			`ABACABA`,
			`3
			1 4
			3 2
			7 1
			`,
		},
		{
			`AAA`,
			`3
			1 3
			2 2
			3 1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF432D, testCases, 0)
}
