//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/371/E
// https://codeforces.com/problemset/status/371/problem/E
func TestCF371E(t *testing.T) {
	t.Log("Current test is [CF371E]")
	testCases := [][2]string{
		{
			`3
			1 100 101
			2`,
			`2 3 `,
		},
	}
	codeforces.AssertEqualStringCase(t, CF371E, testCases, 0)
}
