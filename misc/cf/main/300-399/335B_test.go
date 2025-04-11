//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/335/B
// https://codeforces.com/problemset/status/335/problem/B
func TestCF335B(t *testing.T) {
	t.Log("Current test is [CF335B]")
	testCases := [][2]string{
		{
			`bbbabcbbb`,
			`bbbcbbb`,
		},
		{
			`rquwmzexectvnbanemsmdufrg`,
			`rumenanemur`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF335B, testCases, 0)
}
