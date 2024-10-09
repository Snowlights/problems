//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/628/B
// https://codeforces.com/problemset/status/628/problem/B
func TestCF628B(t *testing.T) {
	t.Log("Current test is [CF628B]")
	testCases := [][2]string{
		{
			`124`,
			`4`,
		},
		{
			`04`,
			`3`,
		},
		{
			`5810438174`,
			`9`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF628B, testCases, 0)
}
