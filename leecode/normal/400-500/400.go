package _00_500

// 409
func longestPalindrome(s string) int {
	h, ans, flag := make(map[byte]int), 0, false
	for _, v := range s {
		h[byte(v)]++
	}
	for _, v := range h {
		ans += v / 2 * 2
		if !flag && v&1 == 1 {
			flag = true
			ans++
		}
	}
	return ans
}

// 413
func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		base := nums[i+1] - nums[i]
		if nums[i+2]-nums[i+1] != base {
			continue
		}
		r := i + 2
		for r < len(nums) && nums[r]-nums[r-1] == base {
			r++
		}
		ans += r - i - 2
	}
	return ans
}

// 416
func canPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 == 1 {
		return false
	}

	k := total >> 1
	dp := make([]int, k+1)
	dp[0] = 1
	for _, v := range nums {
		for j := k; j >= v; j-- {
			dp[j] = dp[j] | dp[j-v]
		}
	}
	return dp[k] == 1
}

// 417
func pacificAtlantic(heights [][]int) [][]int {
	p, a := make([][]bool, len(heights)), make([][]bool, len(heights))
	for i := range p {
		p[i] = make([]bool, len(heights[i]))
		a[i] = make([]bool, len(heights[i]))
	}

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int, r [][]bool)
	dfs = func(i, j int, r [][]bool) {
		r[i][j] = true
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if 0 <= x && x < len(heights) && 0 <= y && y < len(heights[0]) && !r[x][y] &&
				heights[x][y] >= heights[i][j] {
				dfs(x, y, r)
			}
		}
	}

	for i := 0; i < len(heights[0]); i++ {
		dfs(0, i, p)
		dfs(len(heights)-1, i, a)
	}
	for i := 0; i < len(heights); i++ {
		dfs(i, 0, p)
		dfs(i, len(heights[0])-1, a)
	}

	ans := make([][]int, 0)
	for i := range heights {
		for j := range heights[i] {
			if p[i][j] && a[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// 424
func characterReplacement(s string, k int) int {
	h, mm := make(map[byte]int), 0
	left := 0
	for right := range s {
		h[s[right]]++
		mm = max(mm, h[byte(s[right])])
		if right-left+1-mm > k {
			h[s[left]]--
			left++
		}
	}
	return len(s) - left
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
