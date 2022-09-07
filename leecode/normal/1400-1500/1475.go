package _400_1500

// 1480

func runningSum(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i] + nums[i]
	}
	return pre
}

// 1484
// SELECT sell_date,
//       COUNT(DISTINCT product) num_sold,
//       GROUP_CONCAT(DISTINCT product) products
// FROM Activities
// GROUP BY sell_date
