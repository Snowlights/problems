// Code generated by template/gen/leetcode/generator_test.go
package main

import (
	"problems/testutil/leetcode"
	"testing"
)

func Test_d(t *testing.T) {
	t.Log("记得初始化所有全局变量")
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeClassWithFile(t, Constructor, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

// https://leetcode.cn/contest/biweekly-contest-102/problems/design-graph-with-shortest-path-calculator/
// https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/
