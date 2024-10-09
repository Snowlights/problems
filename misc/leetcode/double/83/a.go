package double_83

import "sort"

// 1
func bestHand(ranks []int, suits []byte) string {

	a, b := make(map[int]int), make(map[byte]bool)
	for _, r := range ranks {
		a[r]++
	}
	for _, r := range suits {
		b[r] = true
	}
	if len(b) == 1 {
		return "Flush"
	}
	if len(a) == 5 {
		return "High Card"
	}
	three, two := false, false
	for _, v := range a {
		if v >= 3 {
			three = true
		}
		if v >= 2 {
			two = true
		}
	}
	if three {
		return "Three of a Kind"
	}
	if two {
		return "Pair"
	}
	return ""
}

// 2
func zeroFilledSubarray(nums []int) int64 {
	ans := 0
	pre := 0
	for _, v := range nums {
		if v == 0 {
			ans += pre + 1
			pre++
		} else {
			pre = 0
		}
	}
	return int64(ans)
}

// 3

type NumberContainers struct {
	h   map[int][]int
	idx map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		h:   make(map[int][]int),
		idx: make(map[int]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {

	if this.idx[index] == number {
		return
	}

	val, ok := this.idx[index]
	if ok {
		m, ok := this.h[val]
		if ok && len(m) > 0 {
			idx := sort.Search(len(m), func(i int) bool {
				return m[i] >= index
			})
			m = append(m[:idx], m[idx+1:]...)
			this.h[val] = m
		}
	}

	this.idx[index] = number
	m, ok := this.h[number]
	if !ok {
		m = make([]int, 0)
	}
	if len(m) > 0 {
		idx := sort.Search(len(m), func(i int) bool {
			return m[i] > index
		})
		tmp := append([]int{index}, m[idx:]...)
		m = append(m[:idx], tmp...)
	} else {
		m = append(m, index)
	}
	this.h[number] = m
}

func (this *NumberContainers) Find(number int) int {

	m, ok := this.h[number]
	if !ok {
		return -1
	}

	if len(m) == 0 {
		return -1
	}

	return m[0]
}

// 4
func shortestSequence(rolls []int, k int) int {

	h := make(map[int]bool)
	r := 0
	for _, v := range rolls {
		h[v] = true
		if len(h) == k {
			h = make(map[int]bool)
			r++
		}
	}

	return r + 1
}
