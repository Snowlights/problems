package algorithm

func _() {
	// 位运算 枚举集合
	var s, n int
	// 设元素范围从 0 到 n-1, 挨个判断元素是否在集合 s 中
	for i := 0; i < n; i++ {
		if s>>i&1 == 1 { // i 在 s 中
			// 处理 i 的逻辑
		}
	}

	// 设元素范围从 0 到 n-1 从空集枚举到全集：
	for s := 0; s < 1<<n; s++ {
		// 处理 s 的逻辑
	}

	// 设集合为 s, 从大到小枚举 s 的所有非空子集 sub
	for sub := s; sub > 0; sub = (sub - 1) & s {
		// 处理 sub 的逻辑
	}
}
