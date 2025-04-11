//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1708/B
// https://codeforces.com/problemset/status/1708/problem/B
func TestCF1708B(t *testing.T) {
	t.Log("Current test is [CF1708B]")
	testCases := [][2]string{
		{
			`4
			5 1 5
			9 1000 2000
			10 30 35
			1 1000000000 1000000000`,
			`YES
			1 2 3 4 5
			YES
			1145 1926 1440 1220 1230 1350 1001 1000 1233
			NO
			YES
			1000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1708B, testCases, 0)
}
