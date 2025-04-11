//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/109/C
// https://codeforces.com/problemset/status/109/problem/C
func TestCF109C(t *testing.T) {
	t.Log("Current test is [CF109C]")
	testCases := [][2]string{
		{
			`4
			1 2 4
			3 1 2
			1 4 7`,
			`16`,
		}, {
			`4
			1 2 4
			1 3 47
			1 4 7447`,
			`24`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF109C, testCases, 0)
}
