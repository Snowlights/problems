package main

func distMoney(money, children int) int {
	money -= children // 每人至少 1 美元
	if money < 0 {
		return -1
	}
	ans := min(money/7, children) // 初步分配，让尽量多的人分到 8 美元
	money -= ans * 7
	children -= ans
	if children == 0 && money > 0 || // 必须找一个前面分了 8 美元的人，分完剩余的钱
		children == 1 && money == 3 { // 不能有人恰好分到 4 美元
		ans--
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
