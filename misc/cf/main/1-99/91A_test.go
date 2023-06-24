//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/contest/91/problem/A
// https://codeforces.com/problemset/status/91/problem/A
func TestCF91A(t *testing.T) {
	t.Log("Current test is [CF91A]")
	testCases := [][2]string{
		//{
		//	`abc
		//	xyz`,
		//	`-1`,
		//},
		//{
		//	`abcd
		//	dabc`,
		//	`2`,
		//},
		{
			`ibifgcfdbfdhihbifageaaadegbfbhgeebgdgiafgedchdg
			dedfebcfdigdefdediigcfcafbhhiacgfbeccfchd`,
			`7`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF91A, testCases, 0)
}
