package _00_200

import "strings"

// 161
func reverseWords(s string) string {
	var ans []string
	for _, v := range strings.Split(strings.TrimSpace(s), " ") {
		if v != "" {
			ans = append(ans, v)
		}
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return strings.Join(ans, " ")
}

// 152
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minF = min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 154
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return nums[low]
}

// 167
func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1

	for start < end {
		num := numbers[start] + numbers[end]
		if num == target {
			return []int{start + 1, end + 1}
		}
		if num < target {
			start++
		} else {
			end--
		}
	}

	return nil
}

// 173
type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.inorder(root)
	return
}

func (it *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	it.inorder(node.Left)
	it.arr = append(it.arr, node.Val)
	it.inorder(node.Right)
}

func (it *BSTIterator) Next() int {
	val := it.arr[0]
	it.arr = it.arr[1:]
	return val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.arr) > 0
}

// 174
func maxProfit174(prices []int, fee int) int {
	n := len(prices)
	sell, buy := 0, -prices[0]
	for i := 1; i < n; i++ {
		sell = int(max((sell), (buy + prices[i] - fee)))
		buy = int(max((buy), (sell - prices[i])))
	}
	return sell
}
