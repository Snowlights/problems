package _00_900

import "sort"

// 852
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] > arr[i+1] })
}
