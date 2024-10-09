package main

func resultArray(nums []int) []int {
	a1, a2 := []int{nums[0]}, []int{nums[1]}
	for _, v := range nums[2:] {
		if a1[len(a1)-1] > a2[len(a2)-1] {
			a1 = append(a1, v)
		} else {
			a2 = append(a2, v)
		}
	}
	return append(a1, a2...)
}
