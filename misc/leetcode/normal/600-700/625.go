package _00_700

import (
	"container/heap"
	"math"
)

// 627
// UPDATE salary
// SET
//    sex = CASE sex
//        WHEN 'm' THEN 'f'
//        ELSE 'm'
//    END;
//

// 632
var (
	next  []int
	numsC [][]int
)

func smallestRange(nums [][]int) []int {
	numsC = nums
	rangeLeft, rangeRight := 0, math.MaxInt32
	minRange := rangeRight - rangeLeft
	max := math.MinInt32
	size := len(nums)
	next = make([]int, size)
	h := &IHeap{}
	heap.Init(h)

	for i := 0; i < size; i++ {
		heap.Push(h, i)
		max = Max(max, nums[i][0])
	}

	for {
		minIndex := heap.Pop(h).(int)
		curRange := max - nums[minIndex][next[minIndex]]
		if curRange < minRange {
			minRange = curRange
			rangeLeft, rangeRight = nums[minIndex][next[minIndex]], max
		}
		next[minIndex]++
		if next[minIndex] == len(nums[minIndex]) {
			break
		}
		heap.Push(h, minIndex)
		max = Max(max, nums[minIndex][next[minIndex]])
	}
	return []int{rangeLeft, rangeRight}
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return numsC[h[i]][next[h[i]]] < numsC[h[j]][next[h[j]]] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 633
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		rt := math.Sqrt(float64(c - a*a))
		if rt == math.Floor(rt) {
			return true
		}
	}
	return false
}

// 649
func predictPartyVictory(senate string) string {
	var radiant, dire []int
	n := len(senate)
	for i, v := range senate {
		switch v {
		case 'R':
			radiant = append(radiant, i)
		case 'D':
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		a, b := radiant[0], dire[0]
		radiant, dire = radiant[1:], dire[1:]
		if a < b {
			radiant = append(radiant, a+n)
		} else {
			dire = append(dire, b+n)
		}
	}

	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}
