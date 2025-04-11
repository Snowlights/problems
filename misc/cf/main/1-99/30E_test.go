//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/30/E
// https://codeforces.com/problemset/status/30/problem/E
func TestCF30E(t *testing.T) {
	t.Log("Current test is [CF30E]")
	testCases := [][2]string{
		{
			`abacaba`,
			`1
			1 7`,
		},
		{
			`axbya`,
			`3
			1 1
			2 1
			5 1`,
		},
		{
			`xabyczba`,
			`3
			2 2
			4 1
			7 2`,
		},
		{
			`ntomzzuttrtaapousysvfgelrpqrnljqvfmcyytiheqnjuhpln`,
			`3
			1 1
			9 3
			50 1`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF30E, testCases, 0)
}
