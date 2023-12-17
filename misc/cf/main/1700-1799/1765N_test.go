//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1765/N
// https://codeforces.com/problemset/status/1765/problem/N
func TestCF1765N(t *testing.T) {
	t.Log("Current test is [CF1765N]")
	testCases := [][2]string{
		{
			`5
			10000
			4
			1337
			0
			987654321
			6
			66837494128
			5
			7808652
			3`,
			`1
			1337
			321
			344128
			7052
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1765N, testCases, 0)
}
