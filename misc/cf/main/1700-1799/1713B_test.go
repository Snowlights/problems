//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1713/B
// https://codeforces.com/problemset/status/1713/problem/B
func TestCF1713B(t *testing.T) {
	t.Log("Current test is [CF1713B]")
	testCases := [][2]string{
		{
			`3
			4
			2 3 5 4
			3
			1 2 3
			4
			3 1 3 2`,
			`YES
			YES
			NO`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1713B, testCases, 0)
}
