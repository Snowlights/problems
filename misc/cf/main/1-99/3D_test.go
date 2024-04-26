//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/3/D
// https://codeforces.com/problemset/status/3/problem/D
func TestCF3D(t *testing.T) {
	t.Log("Current test is [CF3D]")
	testCases := [][2]string{
		{
			`(??)
			1 2
			2 8`,
			`4
			()()`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF3D, testCases, 0)
}
