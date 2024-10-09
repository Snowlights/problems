//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1800/F
// https://codeforces.com/problemset/status/1800/problem/F
func TestCF1800F(t *testing.T) {
	t.Log("Current test is [CF1800F]")
	testCases := [][2]string{
		{
			`10
			ftl
			abcdefghijklmnopqrstuvwxy
			abcdeffghijkllmnopqrsttuvwxy
			ffftl
			aabbccddeeffgghhiijjkkllmmnnooppqqrrssttuuvvwwxxyy
			thedevid
			bcdefghhiiiijklmnopqrsuwxyz
			gorillasilverback
			abcdefg
			ijklmnopqrstuvwxyz`,
			`5
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1800F, testCases, 0)
}
