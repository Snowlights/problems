//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/327/A
// https://codeforces.com/problemset/status/327/problem/A
func TestCF327A(t *testing.T) {
	t.Log("Current test is [CF327A]")
	testCases := [][2]string{
		{
			`5
			1 0 0 1 0`,
			`4`,
		},
		{
			`4
			1 0 0 1`,
			`4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF327A, testCases, 0)
}
