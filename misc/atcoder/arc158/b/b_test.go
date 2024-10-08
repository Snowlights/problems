// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc158/submit?taskScreenName=arc158_b
// 对拍：https://atcoder.jp/contests/arc158/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc158_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`4
			-2 -4 4 5`,
			`-0.175000000000000
			-0.025000000000000`,
		},
		{
			`4
			1 1 1 1`,
			`3.000000000000000
			3.000000000000000`,
		},
		{
			`5
			1 2 3 4 5`,
			`0.200000000000000
			1.000000000000000`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/arc158/tasks/arc158_b
