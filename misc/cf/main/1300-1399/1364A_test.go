//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1364/A
// https://codeforces.com/problemset/status/1364/problem/A
func TestCF1364A(t *testing.T) {
	t.Log("Current test is [CF1364A]")
	testCases := [][2]string{
		{
			`3
			3 3
			1 2 3
			3 4
			1 2 3
			2 2
			0 6`,
			`2
			3
			-1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1364A, testCases, 0)
}
