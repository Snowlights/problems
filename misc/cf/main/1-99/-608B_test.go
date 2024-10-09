//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/-608/B
// https://codeforces.com/problemset/status/-608/problem/B
func TestCF-608B(t *testing.T) {
	t.Log("Current test is [CF-608B]")
	testCases := [][2]string{
		{
			``,
			``,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF-608B, testCases, 0)
}
