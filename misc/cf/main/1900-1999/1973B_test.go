//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1973/b
// https://codeforces.com/problemset/status/1973/problem/b
func TestCF1973b(t *testing.T) {
	t.Log("Current test is [CF1973B]")
	testCases := [][2]string{
		{
			`7
			1
			0
			3
			2 2 2
			3
			1 0 2
			5
			3 0 1 4 2
			5
			2 0 4 0 2
			7
			0 0 0 0 1 2 4
			8
			0 1 3 2 2 1 0 3`,
			`1
			1
			3
			4
			4
			7
			3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1973B, testCases, 0)
}
