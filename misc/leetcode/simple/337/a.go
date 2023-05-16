package simple_337

// 1
func evenOddBit(n int) []int {
	ans := make([]int, 2)
	for i := 0; n > 0; i ^= 1 {
		ans[i] += n & 1
		n >>= 1
	}
	return ans
}

// 2
func checkValidGrid(grid [][]int) bool {
	type pair struct{ i, j int }
	pos := make([]pair, len(grid)*len(grid))
	for i, row := range grid {
		for j, x := range row {
			pos[x] = pair{i, j} // 记录坐标
		}
	}
	if pos[0] != (pair{}) { // 必须从左上角出发
		return false
	}
	for i := 1; i < len(pos); i++ {
		dx := abs(pos[i].i - pos[i-1].i)
		dy := abs(pos[i].j - pos[i-1].j)                  // 移动距离
		if (dx != 2 || dy != 1) && (dx != 1 || dy != 2) { // 不合法
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 3
func beautifulSubsets(nums []int, k int) int {
	ans := -1                    // 去掉空集
	cnt := make([]int, 1001+k*2) // 用数组实现比哈希表更快
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans++
			return
		}
		dfs(i + 1)       // 不选
		x := nums[i] + k // 避免负数下标
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++ // 选
			dfs(i + 1)
			cnt[x]-- // 恢复现场
		}
	}
	dfs(0)
	return ans
}

// 4
func findSmallestInteger(nums []int, m int) (mex int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[(x%m+m)%m]++
	}
	for cnt[mex%m] > 0 {
		cnt[mex%m]--
		mex++
	}
	return
}
