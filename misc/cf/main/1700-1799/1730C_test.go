//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1730/C
// https://codeforces.com/problemset/status/1730/problem/C
func TestCF1730C(t *testing.T) {
	t.Log("Current test is [CF1730C]")
	testCases := [][2]string{
		{
			`4
			04829
			9
			01
			314752277691991`,
			`02599
			9
			01
			111334567888999
			`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF1730C, testCases, 0)
}
