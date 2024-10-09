//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/296/B
// https://codeforces.com/problemset/status/296/problem/B
func TestCF296B(t *testing.T) {
	t.Log("Current test is [CF296B]")
	testCases := [][2]string{
		{
			`2
			90
			09
			`,
			`1`,
		},
		{
			`2
			11
			55
			`,
			`0`,
		},
		{
			`5
			?????
			?????
			`,
			`993531194`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF296B, testCases, 0)
}
