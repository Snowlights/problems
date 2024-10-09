//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1187/B
// https://codeforces.com/problemset/status/1187/problem/B
func TestCF1187B(t *testing.T) {
	t.Log("Current test is [CF1187B]")
	testCases := [][2]string{
		{
			`9
			arrayhead
			5
			arya
			harry
			ray
			r
			areahydra`,
			`5
			6
			5
			2
			9`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1187B, testCases, 0)
}
