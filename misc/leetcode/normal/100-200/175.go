package _00_200

// 183
// select Name as Customers from Customers where id not in (
//    select distinct(CustomerId) from Orders
//)

// 187
func findRepeatedDnaSequences(s string) (ans []string) {
	const L = 10
	cnt := map[string]int{}
	for i := 0; i <= len(s)-L; i++ {
		sub := s[i : i+L]
		cnt[sub]++
		if cnt[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return
}

// 189
func rotate(nums []int, k int) {
	k = k % len(nums)
	res := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, res)
}

// 196
// DELETE p1 FROM Person p1,
//    Person p2
// WHERE
//    p1.Email = p2.Email AND p1.Id > p2.Id

// 198
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	// 0是不拿，1是拿
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 199
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := make([]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		ans = append(ans, tmp[len(tmp)-1].Val)
	}

	return ans
}
