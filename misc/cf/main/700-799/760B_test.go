//go:build main
// +build main


package main


import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/760/B
// https://codeforces.com/problemset/status/760/problem/B
func TestCF760B(t *testing.T) {
	t.Log("Current test is [CF760B]")
	testCases := [][2]string{
		{
			`4 6 2`,
			`2`,
		},
		{
			`3 10 3`,
			`4`,
		},
		{
			`3 6 1`,
			`3`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF760B, testCases, 0)
}
