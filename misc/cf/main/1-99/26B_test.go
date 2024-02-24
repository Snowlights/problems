//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/26/B
// https://codeforces.com/problemset/status/26/problem/B
func TestCF26B(t *testing.T) {
	t.Log("Current test is [CF26B]")
	testCases := [][2]string{
		{
			`(()))(`,
			`4`,
		},
		{
			`((()())`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF26B, testCases, 0)
}
