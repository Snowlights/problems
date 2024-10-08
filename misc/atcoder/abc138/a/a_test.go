// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc138/submit?taskScreenName=abc138_a
// 对拍：https://atcoder.jp/contests/abc138/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc138_a&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`3200
			pink`,
			`pink`,
		},
		{
			`3199
			pink`,
			`red`,
		},
		{
			`4049
			red`,
			`red`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc138/tasks/abc138_a
