package main

func reverseDegree(s string) (ans int) {
	for i, c := range s {
		ans += int('{'-c) * (i + 1) // 下标从 1 开始
	}
	return
}
