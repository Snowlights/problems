//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/608/B
// https://codeforces.com/problemset/status/608/problem/B
func TestCF608B(t *testing.T) {
	t.Log("Current test is [CF608B]")
	testCases := [][2]string{
		{
			`01
			00111
			`,
			`3`,
		},
		{
			`0011
			0110
			`,
			`2`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF608B, testCases, 0)
}
