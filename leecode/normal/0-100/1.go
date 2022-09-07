package __100

import (
	"sort"
	"strings"
)

// 1
func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for i, v := range nums {
		if val, ok := h[target-v]; ok {
			return []int{val, i}
		}
		h[v] = i
	}
	return nil
}

// 3
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}

// 11
// 1,8,6,2,5,4,8,3,7
func maxArea(height []int) int {
	//双指针
	length := len(height)
	left, right := 0, length-1
	max := 0
	for left < right {
		width := right - left
		high := 0
		if height[right] > height[left] {
			high = height[left]
			left++
		} else {
			high = height[right]
			right--
		}

		res := width * high
		if res > max {
			max = res
		}
	}
	return max
}

// 窗口内没有重复字符：此时判断i+1与end的关系，超过表示遍历到窗口之外了，增大窗口大小
// 窗口内出现重复字符：此时两个指针都增大index+1，滑动窗口位置到重复字符的后一位
// 遍历结束，返回end-start，窗口大小

// 15
func threeSum(nums []int) [][]int {

	newNums := nums
	if len(newNums) == 3 {
		if newNums[0]+newNums[1]+newNums[2] == 0 {
			return [][]int{newNums}
		}
		return [][]int{}
	}

	sort.Slice(newNums, func(i, j int) bool {
		return newNums[i] < newNums[j]
	})

	res := make([][]int, 0)
	start, end, total := 0, len(newNums)-2, len(newNums)-1

	for start < end {
		if start > 0 && newNums[start] == newNums[start-1] {
			start++
			continue
		}

		left, right := start+1, total

		for left < right {
			tmp := newNums[start] + newNums[left] + newNums[right]
			if tmp == 0 {
				res = append(res, []int{newNums[start], newNums[left], newNums[right]})
				for left < right && newNums[left] == newNums[left+1] {
					left++
				}
				for left < right && newNums[right] == newNums[right-1] {
					right--
				}
			}

			if tmp >= 0 {
				right--
			}

			if tmp < 0 {
				left++
			}

		}

		start++
	}

	return res
}

// 18
func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}

// 20
func isValid(s string) bool {
	st := []byte{}
	for _, b := range s {
		switch b {
		case '{', '(', '[':
			st = append(st, byte(b))
		case '}':
			if len(st) == 0 {
				return false
			}
			v := st[len(st)-1]
			st = st[:len(st)-1]
			if v != '{' {
				return false
			}
		case ']':
			if len(st) == 0 {
				return false
			}
			v := st[len(st)-1]
			st = st[:len(st)-1]
			if v != '[' {
				return false
			}
		case ')':
			if len(st) == 0 {
				return false
			}
			v := st[len(st)-1]
			st = st[:len(st)-1]
			if v != '(' {
				return false
			}
		}

	}
	return len(st) == 0
}

// 21
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	cur := p
	l1, l2 := list1, list2
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}
	for l1 != nil {
		cur.Next = l1
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		cur.Next = l2
		cur = cur.Next
		l2 = l2.Next
	}

	return p.Next
}
