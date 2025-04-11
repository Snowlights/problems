//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2022/B
// https://codeforces.com/problemset/status/2022/problem/B
func TestCF2022B(t *testing.T) {
	t.Log("Current test is [CF2022B]")
	testCases := [][2]string{
		{
			`4
			3 2
			3 1 2
			3 3
			2 1 3
			5 3
			2 2 1 9 2
			7 4
			2 5 3 3 5 2 5`,
			`3
			3
			9
			7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2022B, testCases, 0)
}
