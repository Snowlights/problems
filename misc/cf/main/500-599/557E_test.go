//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/557/E
// https://codeforces.com/problemset/status/557/problem/E
func TestCF557E(t *testing.T) {
	t.Log("Current test is [CF557E]")
	testCases := [][2]string{
		{
			`abbabaab
7`,
			`abaa`,
		},
		{
			`aaaaa
10`,
			`aaa`,
		},
		{
			`bbaabb
13`,
			`bbaabb`,
		}}
	codeforces.AssertEqualStringCase(t, CF557E, testCases, 0)
}
