//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2045/C
// https://codeforces.com/problemset/status/2045/problem/C
func TestCF2045C(t *testing.T) {
	t.Log("Current test is [CF2045C]")
	testCases := [][2]string{
		{
			`sarana
			olahraga`,
			`saga`,
		},
		{
			`berhiber
			wortelhijau`,
			`berhijau`,
		},
		{
			`icpc
			icpc`,
			`icpc`,
		},
		{
			`icpc
			jakarta`,
			`-1`,
		},

	}
	codeforces.AssertEqualStringCase(t, CF2045C, testCases, 0)
}
