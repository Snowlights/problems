//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1081/C
// https://codeforces.com/problemset/status/1081/problem/C
func TestCF1081C(t *testing.T) {
	t.Log("Current test is [CF1081C]")
	testCases := [][2]string{
		{
			`3 3 0`,
			`3`,
		},
		{
			`3 2 1`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1081C, testCases, 0)
}
