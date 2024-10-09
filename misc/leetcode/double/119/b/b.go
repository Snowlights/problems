package main

func removeAlmostEqualCharacters(word string) int {
	i, n, ans := 1, len(word), 0
	for i < n {
		if abs(int(word[i])-int(word[i-1])) <= 1 {
			i += 2
			ans++
		} else {
			i++
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
