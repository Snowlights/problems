package simple_296

import "sort"

// 1
func minMaxGame(nums []int) int {

	res := nums
	for len(res) != 1 {
		newRes := make([]int, 0)
		for i := 0; i < len(res)/2; i++ {
			if i%2 == 0 {
				newRes = append(newRes, min(res[2*i], res[2*i+1]))
			} else {
				newRes = append(newRes, max(res[2*i], res[2*i+1]))
			}
		}
		res = newRes
	}
	return res[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2
func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	ans, idx := 0, 0
	for idx < len(nums) {
		base := nums[idx]
		ans++
		for idx < len(nums) && nums[idx]-base <= k {
			idx++
		}
	}
	return ans
}

// 3
func arrayChange(nums []int, operations [][]int) []int {
	h := make(map[int]int)
	for i, v := range nums {
		h[v] = i
	}
	for _, op := range operations {
		if val, ok := h[op[0]]; ok {
			h[val] = op[1]
			h[op[1]] = val
		}
	}
	return nums
}

// 4
type TextEditor struct {
	s1, s2 []byte
}

func Constructor() TextEditor {
	return TextEditor{}
}

func (this *TextEditor) AddText(text string) {
	this.s1 = append(this.s1, text...)
}

func (this *TextEditor) DeleteText(k int) int {
	k = min(k, len(this.s1))
	this.s1 = this.s1[:len(this.s1)-k]
	return k
}

func (this *TextEditor) CursorLeft(k int) string {
	k = max(len(this.s1)-k, 0)
	tmp := this.s1[k:]
	this.s1 = this.s1[:k]
	this.s2 = append(this.s2, this.rev([]byte(tmp))...)
	return this.text()
}

func (this *TextEditor) CursorRight(k int) string {
	k = max(len(this.s2)-k, 0)
	tmp := this.s2[k:]
	this.s2 = this.s2[:k]
	this.s1 = append(this.s1, this.rev([]byte(tmp))...)
	return this.text()
}

func (this *TextEditor) text() string {
	k := max(0, len(this.s1)-10)
	return string(this.s1[k:])
}

func (this *TextEditor) rev(b []byte) []byte {
	s, e := 0, len(b)-1
	for s <= e {
		b[s], b[e] = b[e], b[s]
		s++
		e--
	}
	return b
}
