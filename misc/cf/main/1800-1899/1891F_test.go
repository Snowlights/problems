//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1891/F
// https://codeforces.com/problemset/status/1891/problem/F
func TestCF1891F(t *testing.T) {
	t.Log("Current test is [CF1891F]")
	testCases := [][2]string{
		{
			`3
			9
			2 1 3
			1 1
			2 2 1
			1 1
			2 3 2
			1 3
			2 1 4
			1 3
			2 3 2
			5
			2 1 1
			1 1
			2 1 -1
			1 1
			2 1 1
			5
			1 1
			1 1
			2 1 1
			2 1 3
			2 2 10`,
			`7 5 8 6 2 
			1 0 1 
			4 14 4`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1891F, testCases, 0)
}
