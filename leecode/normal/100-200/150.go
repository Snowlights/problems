package _00_200

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
