//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/318/B
// https://codeforces.com/problemset/status/318/problem/B
func TestCF318B(t *testing.T) {
	t.Log("Current test is [CF318B]")
	testCases := [][2]string{
		{
			`heavymetalisheavymetal`,
			`3`,
		},
		{
			`heavymetalismetal`,
			`2`,
		},
		{
			`trueheavymetalissotruewellitisalsosoheavythatyoucanalmostfeeltheweightofmetalonyou`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF318B, testCases, 0)
}
