//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1009/B
// https://codeforces.com/problemset/status/1009/problem/B
func TestCF1009B(t *testing.T) {
	t.Log("Current test is [CF1009B]")
	testCases := [][2]string{
		{
			`100210`,
			`001120`,
		},
		{
			`11222121`,
			`11112222`,
		},
		{
			`20`,
			`20`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1009B, testCases, 0)
}
