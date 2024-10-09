//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1984/C2
// https://codeforces.com/problemset/status/1984/problem/C2
func TestCF1984C2(t *testing.T) {
	t.Log("Current test is [CF1984C2]")
	testCases := [][2]string{
		{
			`5
			4
			2 -5 3 -3
			8
			1 4 3 4 1 4 3 4
			3
			-1 -2 -3
			4
			-1000000000 1000000000 1000000000 1000000000
			4
			1 9 8 4`,
			`12
			256
			1
			8
			16`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1984C2, testCases, 0)
}
