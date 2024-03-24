//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1272/F
// https://codeforces.com/problemset/status/1272/problem/F
func TestCF1272F(t *testing.T) {
	t.Log("Current test is [CF1272F]")
	testCases := [][2]string{
		{
			`(())(()
			()))()`,
			`(())()()`,
		},
		//{
		//	`)
		//	((`,
		//	`(())`,
		//},
		//{
		//	`)
		//	)))`,
		//	`((()))`,
		//},
		//{
		//	`())
		//	(()(()(()(`,
		//	`(()()()(()()))`,
		//},
	}
	codeforces.AssertEqualStringCase(t, CF1272FV2, testCases, 0)
}
