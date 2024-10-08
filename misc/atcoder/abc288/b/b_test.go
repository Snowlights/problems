// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc288/submit?taskScreenName=abc288_b
// 对拍：https://atcoder.jp/contests/abc288/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc288_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`5 3
			abc
			aaaaa
			xyz
			a
			def`,
			`aaaaa
			abc
			xyz`,
		},
		{
			`4 4
			z
			zyx
			zzz
			rbg`,
			`rbg
			z
			zyx
			zzz`,
		},
		{
			`3 1
			abc
			arc
			agc`,
			`abc`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc288/tasks/abc288_b
