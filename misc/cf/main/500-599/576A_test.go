//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/576/A
// https://codeforces.com/problemset/status/576/problem/A
func TestCF576A(t *testing.T) {
	t.Log("Current test is [CF576A]")
	testCases := [][2]string{
		{
			`4`,
			`3
			2 4 3 
			`,
		},
		{
			`6`,
			`4
			2 4 3 5 
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF576A, testCases, 0)
}
