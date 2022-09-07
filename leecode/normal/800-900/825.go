package _00_900

// 828
func uniqueLetterString(s string) (ans int) {
	sum, last := 0, [26][2]int{}
	for i := range last {
		last[i] = [2]int{-1, -1}
	}
	for i, c := range s {
		c -= 'A'

		sum += i - last[c][0]*2 + last[c][1]

		ans += sum

		last[c][1] = last[c][0]

		last[c][0] = i
	}
	return
}

// 841
func canVisitAllRooms(rooms [][]int) bool {
	vis := make(map[int]bool)
	vis[0] = true
	q := []int{0}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, vv := range rooms[v] {
				if vis[vv] {
					continue
				}
				q = append(q, vv)
				vis[vv] = true
			}
		}
	}
	return len(rooms) == len(vis)
}
