package _00_300

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// 275
func hIndex(citations []int) int {
	n := len(citations)
	return n - sort.Search(n, func(x int) bool { return citations[x] >= n-x })
}

// 278
func firstBadVersion(n int) int {
	var isBadVersion func(version int) bool
	return sort.Search(n, func(i int) bool {
		return isBadVersion(i) == true
	})
}

// 290
func wordPattern(pattern string, s string) bool {
	ptos, stop := make(map[byte]string), make(map[string]byte)
	parts := strings.Split(s, " ")
	if len(pattern) != len(parts) {
		return false
	}
	for i, p := range pattern {
		sp := parts[i]
		s, ok := ptos[byte(p)]
		if ok && s != sp {
			return false
		}
		pp, ok := stop[sp]
		if ok && pp != byte(p) {
			return false
		}

		stop[sp] = byte(p)
		ptos[byte(p)] = sp
	}

	return true
}

// 297
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var b, e = json.Marshal(root)
	if e != nil {
		panic(e)
	}
	// fmt.Println(string(b))
	return string(b)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var p = new(TreeNode)
	var e = json.Unmarshal([]byte(data), &p)
	if e != nil {
		panic(e)
	}
	//   fmt.Println(p)
	return p
}

// 299
func getHint(secret string, guess string) string {
	right, wrong := 0, 0
	sh, gh := make(map[byte]int), make(map[byte]int)
	for i := range secret {
		if secret[i] == guess[i] {
			right++
			continue
		}
		sh[secret[i]]++
		gh[guess[i]]++
	}
	for k, v := range sh {
		wrong += min(v, gh[k])
	}

	return fmt.Sprintf("%dA%dB", right, wrong)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
