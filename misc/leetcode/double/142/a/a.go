package main

func possibleStringCount(word string) int {
	ans := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			ans++
		}
	}
	return ans
}
