//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/804/B
// https://codeforces.com/problemset/status/804/problem/B
func TestCF804B(t *testing.T) {
	t.Log("Current test is [CF804B]")
	testCases := [][2]string{
		{
			`ab`,
			`1`,
		},
		{
			`aab`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF804B, testCases, 0)
}
