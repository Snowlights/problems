//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1957/D
// https://codeforces.com/problemset/status/1957/problem/D
func TestCF1957D(t *testing.T) {
	t.Log("Current test is [CF1957D]")
	testCases := [][2]string{
		{
			`3
			3
			6 2 4
			1
			3
			5
			7 3 7 2 1`,
			`4
			0
			16`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1957D, testCases, 0)
}
