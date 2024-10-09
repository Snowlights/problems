//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/282/E
// https://codeforces.com/problemset/status/282/problem/E
func TestCF282E(t *testing.T) {
	t.Log("Current test is [CF282E]")
	testCases := [][2]string{
		{
			`2
			1 2`,
			`3`,
		},
		{
			`3
			1 2 3`,
			`3`,
		},
		{
			`2
			1000 1000`,
			`1000`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF282E, testCases, 0)
}
