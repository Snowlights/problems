//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1791/F
// https://codeforces.com/problemset/status/1791/problem/F
func TestCF1791F(t *testing.T) {
	t.Log("Current test is [CF1791F]")
	testCases := [][2]string{
		{
			`3
			5 8
			1 420 69 1434 2023
			1 2 3
			2 2
			2 3
			2 4
			1 2 5
			2 1
			2 3
			2 5
			2 3
			9999 1000
			1 1 2
			2 1
			2 2
			1 1
			1
			2 1`,
			`6
			15
			1434
			1
			6
			7
			36
			1
			1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1791F, testCases, 0)
}
