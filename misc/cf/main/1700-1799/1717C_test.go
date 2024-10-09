//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1717/C
// https://codeforces.com/problemset/status/1717/problem/C
func TestCF1717C(t *testing.T) {
	t.Log("Current test is [CF1717C]")
	testCases := [][2]string{
		{
			`5
			3
			1 2 5
			1 2 5
			2
			2 2
			1 3
			4
			3 4 1 2
			6 4 2 5
			3
			2 4 1
			4 5 3
			5
			1 2 3 4 5
			6 5 6 7 6`,
			`YES
			NO
			NO
			NO
			YES
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1717C, testCases, 0)
}
