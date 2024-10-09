package b

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestRampartDefensiveLine(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, rampartDefensiveLine, "b.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
