//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/292/D
// https://codeforces.com/problemset/status/292/problem/D
func TestCF292D(t *testing.T) {
	t.Log("Current test is [CF292D]")
	testCases := [][2]string{
		{
			`6 5
			1 2
			5 4
			2 3
			3 1
			3 6
			6
			1 3
			2 5
			1 5
			5 5
			2 4
			3 3
			`,
			`4
			5
			6
			3
			4
			2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF292D, testCases, 0)
}
