//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1400/G
// https://codeforces.com/problemset/status/1400/problem/G
func TestCF1400G(t *testing.T) {
	t.Log("Current test is [CF1400G]")
	testCases := [][2]string{
		{
			`3 0
			1 1
			2 3
			1 3`,
			`3`,
		},
		{
			`3 1
			1 1
			2 3
			1 3
			2 3`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1400G, testCases, 0)
}
