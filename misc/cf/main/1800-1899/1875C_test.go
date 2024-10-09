//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1875/C
// https://codeforces.com/problemset/status/1875/problem/C
func TestCF1875C(t *testing.T) {
	t.Log("Current test is [CF1875C]")
	testCases := [][2]string{
		{
			`4
			10 5
			1 5
			10 4
			3 4`,
			`0
			-1
			2
			5`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1875C, testCases, 0)
}
