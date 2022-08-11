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
