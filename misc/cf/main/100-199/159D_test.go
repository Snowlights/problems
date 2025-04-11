//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/159/D
// https://codeforces.com/problemset/status/159/problem/D
func TestCF159D(t *testing.T) {
	t.Log("Current test is [CF159D]")
	testCases := [][2]string{
		{
			`aa`,
			`1`,
		},
		{
			`aaa`,
			`5`,
		},
		{
			`abacaba`,
			`36`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF159D, testCases, 0)
}
