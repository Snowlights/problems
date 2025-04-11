//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1881/D
// https://codeforces.com/problemset/status/1881/problem/D
func TestCF1881D(t *testing.T) {
	t.Log("Current test is [CF1881D]")
	testCases := [][2]string{
		{
			`7
			5
			100 2 50 10 1
			3
			1 1 1
			4
			8 2 4 2
			4
			30 50 27 20
			2
			75 40
			2
			4 4
			3
			2 3 1`,
			`YES
			YES
			NO
			YES
			NO
			YES
			NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1881D, testCases, 0)
}
