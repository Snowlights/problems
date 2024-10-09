package d

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestEvolutionaryRecord(t *testing.T) {
	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, evolutionaryRecord, "d.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
