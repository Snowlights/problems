//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1926/G
// https://codeforces.com/problemset/status/1926/problem/G
func TestCF1926G(t *testing.T) {
	t.Log("Current test is [CF1926G]")
	testCases := [][2]string{
		{
			`3
			3
			1 1
			CSP
			4
			1 2 2
			PCSS
			4
			1 2 2
			PPSS`,
			`1
			1
			2
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1926G, testCases, 0)
}
