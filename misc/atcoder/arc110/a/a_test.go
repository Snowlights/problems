// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc110/submit?taskScreenName=arc110_a
// 对拍：https://atcoder.jp/contests/arc110/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc110_a&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`3`,
			`7`,
		},
		{
			`10`,
			`39916801`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/arc110/tasks/arc110_a
