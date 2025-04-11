//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1687/A
// https://codeforces.com/problemset/status/1687/problem/A
func TestCF1687A(t *testing.T) {
	t.Log("Current test is [CF1687A]")
	testCases := [][2]string{
		{
			`4
			5 2
			5 6 1 2 3
			5 7
			5 6 1 2 3
			1 2
			999999
			5 70000
			1000000000 1000000000 1000000000 1000000000 1000000000`,
			`12
			37
			1000000
			5000349985`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1687A, testCases, 0)
}
