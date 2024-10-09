//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1946/C
// https://codeforces.com/problemset/status/1946/problem/C
func TestCF1946C(t *testing.T) {
	t.Log("Current test is [CF1946C]")
	testCases := [][2]string{
		{
			`6
			5 1
			1 2
			1 3
			3 4
			3 5
			2 1
			1 2
			6 1
			1 2
			2 3
			3 4
			4 5
			5 6
			3 1
			1 2
			1 3
			8 2
			1 2
			1 3
			2 4
			2 5
			3 6
			3 7
			3 8
			6 2
			1 2
			2 3
			1 4
			4 5
			5 6`,
			`2
			1
			3
			1
			1
			2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1946C, testCases, 0)
}
