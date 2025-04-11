//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1955/D
// https://codeforces.com/problemset/status/1955/problem/D
func TestCF1955D(t *testing.T) {
	t.Log("Current test is [CF1955D]")
	testCases := [][2]string{
		{
			`5
			7 4 2
			4 1 2 3 4 5 6
			1 2 3 4
			7 4 3
			4 1 2 3 4 5 6
			1 2 3 4
			7 4 4
			4 1 2 3 4 5 6
			1 2 3 4
			11 5 3
			9 9 2 2 10 9 7 6 3 6 3
			6 9 7 8 10
			4 1 1
			4 1 5 6
			6`,
			`4
			3
			2
			4
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1955D, testCases, 0)
}
