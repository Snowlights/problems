package f

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestTreeOfInfiniteSouls(t *testing.T) {

	// todo undone 存在溢出问题
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, treeOfInfiniteSouls, "f.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
